package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/just-install/just-install-updater-go/jiup"
	"github.com/just-install/just-install-updater-go/jiup/registry"
	"github.com/ogier/pflag"
)

func main() {
	verbose := pflag.BoolP("verbose", "v", false, "Show more output")
	dryRun := pflag.BoolP("dry-run", "d", false, "Do not actually write the changes")
	// testLinks := pflag.BoolP("test-links", "t", false, "Test download links for updated entries")
	quiet := pflag.BoolP("quiet", "q", false, "Do not output progress info")
	help := pflag.Bool("help", false, "Show this help text")
	pflag.Parse()

	if *help || pflag.NArg() < 1 {
		helpExit()
	}

	registryPath := pflag.Arg(0)
	buf, err := ioutil.ReadFile(registryPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening registry: %v\n", err)
		os.Exit(1)
	}
	r, err := registry.NewFromJSON(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing registry: %v\n", err)
		os.Exit(1)
	}

	u := jiup.New(r)
	if pflag.NArg() > 1 {
		u, err = jiup.NewForPackages(r, pflag.Args()[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	updated, unchanged, norule, skipped, errored := u.Update(!*quiet, *verbose)

	if !*dryRun {
		bufn, err := u.Registry.GetJSON()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating new JSON: %v\n", err)
			os.Exit(1)
		}

		err = ioutil.WriteFile(registryPath, bufn, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing new registry: %v\n", err)
			os.Exit(1)
		}
	}

	// test dl links for updated if flag set

	fmt.Printf("\nSummary: %d updated, %d unchanged, %d norule, %d skipped, %d errored\n\n", len(updated), len(unchanged), len(norule), len(skipped), len(errored))
	if len(norule) > 0 {
		fmt.Printf("No rule:\n")
		for _, pkgName := range norule {
			fmt.Printf("  %s\n", pkgName)
		}
	}
	if len(unchanged) > 0 {
		fmt.Printf("Unchanged:\n")
		for _, pkgName := range unchanged {
			fmt.Printf("  %s\n", pkgName)
		}
	}
	if len(updated) > 0 {
		fmt.Printf("Updated:\n")
		for pkgName, version := range updated {
			fmt.Printf("  %s to %s\n", pkgName, version)
		}
	}
	if len(errored) > 0 {
		fmt.Printf("Errors:\n")
		for pkgName, err := range errored {
			fmt.Fprintf(os.Stderr, "  %s: %v\n", pkgName, err)
		}
	}
	fmt.Printf("\nSummary: %d updated, %d unchanged, %d norule, %d skipped, %d errored\n", len(updated), len(unchanged), len(norule), len(skipped), len(errored))

	if *dryRun {
		fmt.Printf("\nDRY RUN. NO CHANGES WERE MADE.\n")
	}
}

func helpExit() {
	fmt.Fprintf(os.Stderr, "Usage: just-install-updater [options] registry [packages...]\n\n")
	pflag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nArguments:\n  registry is the path to just-install.json\n  packages are the packages to update (default is all)\n")
	os.Exit(1)
}

func errExit() {
	os.Exit(1)
}
