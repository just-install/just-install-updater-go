package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/just-install/just-install-updater-go/jiup/rules"
	"github.com/spf13/pflag"
)

// KnownBroken contains packages known to be broken.
// Errors during the check will not count as a failure.
var KnownBroken = []string{}

const overwrite = "\x1b[1A\x1b[K\r      \r" // up 1, clear line, carriage return (fallback for when the escapes aren't supported)

func main() {
	//verbose := pflag.BoolP("verbose", "v", false, "Show more output")
	nodownload := pflag.BoolP("no-download", "d", false, "Do not test downloadability")
	downloadLinks := pflag.BoolP("download-links", "l", false, "Show download links")
	writeBroken := pflag.StringP("write-broken", "b", "", "If set, broken rules will be written into the specified file")
	help := pflag.Bool("help", false, "Show this help text")
	pflag.Parse()

	if *help {
		helpExit()
	}

	working, broken, knownBroken := testAll(*nodownload, *downloadLinks, pflag.Args())

	fmt.Printf("\nSummary: %d working, %d broken, %d known broken\n", len(working), len(broken), len(knownBroken))

	if *writeBroken != "" {
		arr := []string{}
		for _, x := range knownBroken {
			arr = append(arr, x+":"+"manually marked as broken")
		}
		for x, err := range broken {
			arr = append(arr, x+":"+strings.ReplaceAll(err.Error(), "\n", " "))
		}
		sort.Strings(arr)

		if err := ioutil.WriteFile(*writeBroken, []byte(strings.Join(arr, "\n")+"\n"), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to write broken rules to %q: %v\n", *writeBroken, err)
			os.Exit(1)
		}
	}

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

		fmt.Printf("    %s: testing\n", p)

		c := false
		for _, kb := range KnownBroken {
			if p == kb {
				knownBroken = append(knownBroken, p)
				fmt.Printf("%s -  %s: manually marked as broken\n", overwrite, p)
				c = true
				break
			}
		}
		if c {
			continue
		}

		version, err := vfn()
		if err != nil {
			fmt.Printf("%s ✗  %s: %v\n", overwrite, p, err)
			if strings.Contains(err.Error(), "Client.Timeout") {
				fmt.Printf(" [IGNORING TIMEOUT]")
			} else {
				broken[p] = err
			}
			continue
		}
		if strings.TrimSpace(version) == "" {
			broken[p] = errors.New("empty version")
			fmt.Printf("%s ✗  %s: %v\n", overwrite, p, broken[p])
			continue
		}
		if strings.TrimSpace(version) != version {
			broken[p] = errors.New("version has whitespace (probably a bad regexp)")
			fmt.Printf("%s ✗  %s: %v\n", overwrite, p, broken[p])
			continue
		}
		if strings.HasSuffix(version, ".") {
			broken[p] = errors.New("version ends with a dot (probably a bad regexp)")
			fmt.Printf("%s ✗  %s: %v\n", overwrite, p, broken[p])
			continue
		}

		x86dl, x64dl, err := dfn(version)
		if err != nil {
			fmt.Printf("%s ✗  %s: %v\n", overwrite, p, err)
			if strings.Contains(err.Error(), "Client.Timeout") {
				fmt.Print(" [IGNORING TIMEOUT]")
			} else {
				broken[p] = err
			}
			continue
		}
		if x86dl == nil && x64dl == nil {
			broken[p] = errors.New("one or both of x86 and x64 must be defined")
			fmt.Printf("%s ✗  %s: %v\n", overwrite, p, broken[p])
			continue
		}
		if (x86dl != nil && strings.TrimSpace(*x86dl) == "") || (x64dl != nil && strings.TrimSpace(*x64dl) == "") {
			broken[p] = errors.New("use nil if no link, not a blank string")
			fmt.Printf("%s ✗  %s: %v\n", overwrite, p, broken[p])
			continue
		}

		res := fmt.Sprintf(" ✓  %s: %s", p, version)
		for _, l := range []struct {
			arch string
			link *string
		}{{"x86", x86dl}, {"x86_64", x64dl}} {
			if l.link == nil {
				continue
			}
			if !strings.HasPrefix(*l.link, "http") {
				broken[p] = fmt.Errorf("%s link (%s) does not start with http", l.arch, *l.link)
				fmt.Printf("%s ✗  %s: %v\n", overwrite, p, broken[p])
				break
			}
			if !nodownload {
				code, mime, err := testDL(*l.link)
				if err != nil && !(p == "tightvnc" && strings.Contains(err.Error(), "connection reset")) {
					fmt.Printf("%s ✗  %s: %v\n", overwrite, p, err)
					if strings.Contains(err.Error(), "Client.Timeout") {
						fmt.Print(" [IGNORING TIMEOUT]")
					} else {
						broken[p] = err
					}
					break
				}
				if code != 200 {
					broken[p] = fmt.Errorf("%s download status code %d (%s)", l.arch, code, *l.link)
					fmt.Printf("%s ✗  %s: %v\n", overwrite, p, broken[p])
					break
				}
				if strings.HasPrefix(mime, "text/html") && !strings.Contains(*l.link, "sourceforge") && !strings.Contains(*l.link, "freefilesync") {
					broken[p] = fmt.Errorf("%s download mime text/html", l.arch)
					fmt.Printf("%s ✗  %s: %v (%s)\n", overwrite, p, broken[p], *l.link)
					break
				}
			}
			if downloadLinks {
				res = fmt.Sprintf("%s %s(%s)", res, l.arch, *l.link)
			} else {
				res += "                        " // workaround for carriage returns
			}
		}
		if broken[p] == nil {
			working = append(working, p)
			fmt.Printf("%s%s\n", overwrite, res)
		}
	}

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
