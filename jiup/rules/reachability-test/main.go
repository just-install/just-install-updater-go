package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/just-install/just-install-updater-go/jiup/rules"
	"github.com/ogier/pflag"
)

// KnownBroken contains packages known to be broken.
// Errors during the check will not count as a failure.
var KnownBroken = []string{
	"freefilesync",  // The server is unreliable.
	"octave",        // The server is unreliable.
	"jre",           // The server is unreliable.
	"gimp",          // The tests fail, but it seems to work fine when manually testing it
	"audacity",      // https://github.com/just-install/just-install-updater-go/issues/17
	"classic-shell", // https://github.com/just-install/just-install-updater-go/issues/17
	"qbittorrent",   // https://github.com/just-install/just-install-updater-go/issues/17
}

func main() {
	//verbose := pflag.BoolP("verbose", "v", false, "Show more output")
	nodownload := pflag.BoolP("no-download", "d", false, "Do not test downloadability")
	downloadLinks := pflag.BoolP("download-links", "l", false, "Show download links")
	help := pflag.Bool("help", false, "Show this help text")
	pflag.Parse()

	if *help {
		helpExit()
	}

	working, broken, knownBroken := testAll(*nodownload, *downloadLinks, pflag.Args())

	fmt.Printf("\nSummary: %d working, %d broken, %d known broken\n", len(working), len(broken), len(knownBroken))

	if len(broken) > 0 {
		os.Exit(1)
	}

	os.Exit(0)
}

func helpExit() {
	fmt.Fprintf(os.Stderr, "Usage: reachability-test [options] [packages...]\n\n")
	pflag.PrintDefaults()
	os.Exit(1)
}

func testAll(nodownload, downloadLinks bool, packages []string) ([]string, map[string]error, []string) {
	working := []string{}
	broken := map[string]error{}
	knownBroken := []string{}
	// TODO: multithreaded for loop

	allrules := []string{}
	for p := range rules.GetRules() {
		allrules = append(allrules, p)
	}
	sort.Strings(allrules)

	for _, p := range allrules {
		vfn, dfn, ok := rules.GetRule(p)
		if !ok {
			panic("could not get rule which should exist: " + p)
		}

		if len(packages) != 0 {
			c := true
			for _, pp := range packages {
				if pp == p {
					c = false
				}
			}
			if c {
				continue
			}
		}

		fmt.Printf("\n    %s: testing", p)

		c := false
		for _, kb := range KnownBroken {
			if p == kb {
				knownBroken = append(knownBroken, p)
				fmt.Printf("\r -  %s: manually marked as broken", p)
				c = true
				break
			}
		}
		if c {
			continue
		}

		version, err := vfn()
		if err != nil {
			broken[p] = err
			fmt.Printf("\r ✗  %s: %v", p, broken[p])
			continue
		}
		if strings.TrimSpace(version) == "" {
			broken[p] = errors.New("empty version")
			fmt.Printf("\r ✗  %s: %v", p, broken[p])
			continue
		}
		if strings.TrimSpace(version) != version {
			broken[p] = errors.New("version has whitespace (probably a bad regexp)")
			fmt.Printf("\r ✗  %s: %v", p, broken[p])
			continue
		}
		if strings.HasSuffix(version, ".") {
			broken[p] = errors.New("version ends with a dot (probably a bad regexp)")
			fmt.Printf("\r ✗  %s: %v", p, broken[p])
			continue
		}

		x86dl, x64dl, err := dfn(version)
		if err != nil {
			broken[p] = err
			fmt.Printf("\r ✗  %s: %v", p, broken[p])
			continue
		}
		if strings.TrimSpace(x86dl) == "" {
			broken[p] = errors.New("empty x86 download link")
			fmt.Printf("\r ✗  %s: %v", p, broken[p])
			continue
		}
		if !strings.HasPrefix(x86dl, "http") {
			broken[p] = fmt.Errorf("x86 link (%s) does not start with http", x86dl)
			fmt.Printf("\r ✗  %s: %v", p, broken[p])
			continue
		}
		if !nodownload {
			code, mime, err := testDL(x86dl)
			if err != nil && !(p == "tightvnc" && strings.Contains(err.Error(), "connection reset")) {
				broken[p] = err
				fmt.Printf("\r ✗  %s: %v", p, broken[p])
				continue
			}
			if code != 200 {
				broken[p] = fmt.Errorf("x86 download status code %d", code)
				fmt.Printf("\r ✗  %s: %v", p, broken[p])
				continue
			}
			if strings.HasPrefix(mime, "text/html") && !strings.Contains(x86dl, "sourceforge") && !strings.Contains(x86dl, "oracle") {
				broken[p] = errors.New("x86 download mime text/html")
				fmt.Printf("\r ✗  %s: %v", p, broken[p])
				continue
			}
		}
		if x64dl != nil {
			if strings.TrimSpace(*x64dl) == "" {
				broken[p] = errors.New("empty x86_64 download link")
				fmt.Printf("\r ✗  %s: %v", p, broken[p])
				continue
			}
			if !strings.HasPrefix(*x64dl, "http") {
				broken[p] = fmt.Errorf("x86_64 link (%s) does not start with http", *x64dl)
				fmt.Printf("\r ✗  %s: %v", p, broken[p])
				continue
			}
			if !nodownload {
				code, mime, err := testDL(*x64dl)
				if err != nil && !(p == "tightvnc" && strings.Contains(err.Error(), "connection reset")) {
					broken[p] = err
					fmt.Printf("\r ✗  %s: %v", p, broken[p])
					continue
				}
				if code != 200 {
					broken[p] = fmt.Errorf("x86_64 download status code %d", code)
					fmt.Printf("\r ✗  %s: %v", p, broken[p])
					continue
				}
				if strings.HasPrefix(mime, "text/html") && !strings.Contains(*x64dl, "sourceforge") && !strings.Contains(*x64dl, "oracle") {
					broken[p] = errors.New("x86_64 download mime text/html")
					fmt.Printf("\r ✗  %s: %v", p, broken[p])
					continue
				}
			}
		}

		working = append(working, p)
		if downloadLinks {
			if x64dl == nil {
				fmt.Printf("\r ✓  %s: %s x86(%s)", p, version, x86dl)
			} else {
				fmt.Printf("\r ✓  %s: %s x86(%s) x64(%s)", p, version, x86dl, *x64dl)
			}
		} else {
			fmt.Printf("\r ✓  %s: %s           ", p, version)
		}
	}
	fmt.Printf("\n")

	return working, broken, knownBroken
}

func testDL(url string) (code int, mime string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		resp, err = http.Head(url)
		if err != nil {
			return 0, "", err
		}
		defer resp.Body.Close()
	}
	defer resp.Body.Close()

	code = resp.StatusCode
	mime = resp.Header.Get("Content-Type")

	return code, mime, nil
}
