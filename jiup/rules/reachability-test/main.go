package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/just-install/just-install-updater-go/jiup/rules"
	"github.com/ogier/pflag"
)

func main() {
	//verbose := pflag.BoolP("verbose", "v", false, "Show more output")
	nodownload := pflag.BoolP("no-download", "d", false, "Do not test downloadability")
	help := pflag.Bool("help", false, "Show this help text")
	pflag.Parse()

	if *help || pflag.NArg() != 0 {
		helpExit()
	}

	working, broken := testAll(*nodownload)

	fmt.Printf("\nSummary: %d working, %d broken\n", len(working), len(broken))

	if len(broken) > 0 {
		os.Exit(1)
	}

	os.Exit(0)
}

func helpExit() {
	fmt.Fprintf(os.Stderr, "Usage: reachability-test [options]\n\n")
	pflag.PrintDefaults()
	os.Exit(1)
}

func testAll(nodownload bool) ([]string, map[string]error) {
	working := []string{}
	broken := map[string]error{}
	// TODO: multithreaded for loop
	for p, r := range rules.GetRules() {
		fmt.Printf("\n    %s: testing", p)

		version, err := r.V()
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

		x86dl, x64dl, err := r.D(version)
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
			if err != nil {
				broken[p] = err
				fmt.Printf("\r ✗  %s: %v", p, broken[p])
				continue
			}
			if code != 200 {
				broken[p] = fmt.Errorf("x86 download status code %d", code)
				fmt.Printf("\r ✗  %s: %v", p, broken[p])
				continue
			}
			if strings.HasPrefix(mime, "text/html") {
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
				if err != nil {
					broken[p] = err
					fmt.Printf("\r ✗  %s: %v", p, broken[p])
					continue
				}
				if code != 200 {
					broken[p] = fmt.Errorf("x86_64 download status code %d", code)
					fmt.Printf("\r ✗  %s: %v", p, broken[p])
					continue
				}
				if strings.HasPrefix(mime, "text/html") {
					broken[p] = errors.New("x86_64 download mime text/html")
					fmt.Printf("\r ✗  %s: %v", p, broken[p])
					continue
				}
			}
		}

		working = append(working, p)
		fmt.Printf("\r ✓  %s: %s           ", p, version)
	}
	fmt.Printf("\n")

	return working, broken
}

func testDL(url string) (code int, mime string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, "", err
	}

	code = resp.StatusCode
	mime = resp.Header.Get("Content-Type")

	resp.Body.Close()

	return code, mime, nil
}
