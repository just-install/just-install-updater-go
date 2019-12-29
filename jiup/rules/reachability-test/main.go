package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/just-install/just-install-updater-go/jiup/rules"
	"github.com/spf13/pflag"
)

// KnownBroken contains packages known to be broken.
// Errors during the check will not count as a failure.
var KnownBroken = []string{
	"audacity",         // https://github.com/just-install/just-install-updater-go/issues/17
	"cryptomator",      // https://github.com/just-install/just-install-updater-go/issues/15
	"jdk",              // The server is unreliable.
	"crystaldisk-info", // TODO: fixme
	"ditto",            // Temporary server issues
	"mumble",           // latest rc doesn't have windows builds
	"deluge",           // Version 2.0.x doesn't have Windows builds
	"smplayer",         // SMPlayer 19.10.2 is a source-only download
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
			fmt.Printf("\r ✗  %s: %v", p, err)
			if strings.Contains(err.Error(), "Client.Timeout") {
				fmt.Printf(" [IGNORING TIMEOUT]")
			} else {
				broken[p] = err
			}
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
			fmt.Printf("\r ✗  %s: %v", p, err)
			if strings.Contains(err.Error(), "Client.Timeout") {
				fmt.Print(" [IGNORING TIMEOUT]")
			} else {
				broken[p] = err
			}
			continue
		}
		if x86dl == nil && x64dl == nil {
			broken[p] = errors.New("one or both of x86 and x64 must be defined")
			fmt.Printf("\r ✗  %s: %v", p, broken[p])
			continue
		}
		if (x86dl != nil && strings.TrimSpace(*x86dl) == "") || (x64dl != nil && strings.TrimSpace(*x64dl) == "") {
			broken[p] = errors.New("use nil if no link, not a blank string")
			fmt.Printf("\r ✗  %s: %v", p, broken[p])
			continue
		}

		res := fmt.Sprintf("\r ✓  %s: %s", p, version)
		for _, l := range []struct {
			arch string
			link *string
		}{{"x86", x86dl}, {"x86_64", x64dl}} {
			if l.link == nil {
				continue
			}
			if !strings.HasPrefix(*l.link, "http") {
				broken[p] = fmt.Errorf("%s link (%s) does not start with http", l.arch, *l.link)
				fmt.Printf("\r ✗  %s: %v", p, broken[p])
				continue
			}
			if !nodownload {
				code, mime, err := testDL(*l.link)
				if err != nil && !(p == "tightvnc" && strings.Contains(err.Error(), "connection reset")) {
					fmt.Printf("\r ✗  %s: %v", p, err)
					if strings.Contains(err.Error(), "Client.Timeout") {
						fmt.Print(" [IGNORING TIMEOUT]")
					} else {
						broken[p] = err
					}
					continue
				}
				if code != 200 {
					broken[p] = fmt.Errorf("%s download status code %d", l.arch, code)
					fmt.Printf("\r ✗  %s: %v", p, broken[p])
					continue
				}
				if strings.HasPrefix(mime, "text/html") && !strings.Contains(*l.link, "sourceforge") && !strings.Contains(*l.link, "freefilesync") {
					broken[p] = fmt.Errorf("%s download mime text/html", l.arch)
					fmt.Printf("\r ✗  %s: %v", p, broken[p])
					continue
				}
			}
			if downloadLinks {
				res = fmt.Sprintf("%s %s(%s)", res, l.arch, *l.link)
			} else {
				res += "                " // workaround for carriage returns
			}
		}
		working = append(working, p)
		fmt.Print(res)
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
