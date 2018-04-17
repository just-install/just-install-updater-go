package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/just-install/just-install-updater-go/jiup"
	"github.com/just-install/just-install-updater-go/jiup/registry"
	"github.com/ogier/pflag"
)

func main() {
	verbose := pflag.BoolP("verbose", "v", false, "Show more output")
	dryRun := pflag.BoolP("dry-run", "d", false, "Do not actually write the changes")
	force := pflag.BoolP("force", "f", false, "Update all entries including ones with a matching version")
	commitMessageFile := pflag.StringP("commit-message-file", "c", "", "If set, jiup-go will save a commit message describing the changes to a file.")
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

	updated, unchanged, norule, skipped, errored := u.Update(!*quiet, *verbose, *force)

	if commitMessageFile != nil && *commitMessageFile != "" {
		pkgs := []string{}
		pkgvs := []string{}
		for pkg, v := range updated {
			pkgs = append(pkgs, pkg)
			pkgvs = append(pkgvs, "  - "+pkg+" ("+v+")")
		}

		cMessage := "jiup-go automatic commit"
		if len(pkgs) > 0 {
			cMessage = cMessage + ": updated " + listify(pkgs)
		}

		cMessage = cMessage + fmt.Sprintf("\n\n%d updated, %d unchanged, %d norule (%.0f%%), %d skipped, %d errored\n", len(updated), len(unchanged), len(norule), float32(len(norule))/float32(len(u.Registry.Packages))*100.0, len(skipped), len(errored))

		if len(updated) > 0 {
			cMessage = cMessage + "\nUpdated:\n" + strings.Join(pkgvs, "\n") + "\n"
		}
		if len(errored) > 0 {
			errs := []string{}
			for pkg, err := range errored {
				errs = append(errs, "  - "+pkg+" ("+err.Error()+")")
			}
			cMessage = cMessage + "\nErrors:\n" + strings.Join(errs, "\n") + "\n"
		}

		err := ioutil.WriteFile(*commitMessageFile, []byte(cMessage), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing commit message: %v\n", err)
			os.Exit(1)
		}
	}

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

	fmt.Printf("\n===== RESULTS =====\n")
	if len(norule) > 0 {
		fmt.Printf("No rule:\n")
		for _, pkgName := range norule {
			fmt.Printf("  %s\n", pkgName)
		}
		fmt.Printf("\n")
	}
	if len(unchanged) > 0 {
		fmt.Printf("Unchanged:\n")
		for _, pkgName := range unchanged {
			fmt.Printf("  %s\n", pkgName)
		}
		fmt.Printf("\n")
	}
	if len(updated) > 0 {
		fmt.Printf("Updated:\n")
		for pkgName, version := range updated {
			fmt.Printf("  %s to %s\n", pkgName, version)
		}
		fmt.Printf("\n")
	}
	if len(errored) > 0 {
		fmt.Printf("Errors:\n")
		for pkgName, err := range errored {
			fmt.Fprintf(os.Stderr, "  %s: %v\n", pkgName, err)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Summary: %d updated, %d unchanged, %d norule (%.0f%%), %d skipped, %d errored\n", len(updated), len(unchanged), len(norule), float32(len(norule))/float32(len(u.Registry.Packages))*100.0, len(skipped), len(errored))

	if *dryRun {
		fmt.Printf("\nDRY RUN. NO CHANGES WERE MADE.\n")
	}

	os.Exit(0)
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

func listify(arr []string) string {
	switch len(arr) {
	case 0:
		return ""
	case 1:
		return arr[0]
	case 2:
		return arr[0] + " and " + arr[1]
	default:
		return strings.Join(arr[:len(arr)-1], ", ") + ", and " + arr[len(arr)-1]
	}
}
