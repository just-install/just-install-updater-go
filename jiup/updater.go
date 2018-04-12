package jiup

import (
	"errors"
	"fmt"

	"github.com/just-install/just-install-updater-go/jiup/rules"
	"github.com/just-install/just-install-updater-go/jiup/rules/helpers"

	"github.com/just-install/just-install-updater-go/jiup/registry"
)

// Updater represents an instance of the updater.
type Updater struct {
	Registry *registry.Registry
	packages []string
}

// ErrNoSuchPackage is returned if one or more specified packages does not exist.
var ErrNoSuchPackage = errors.New("one or more of the specified packages does not exist")

// New returns a new instance of Updater.
func New(registry *registry.Registry) *Updater {
	return &Updater{
		Registry: registry,
		packages: []string{},
	}
}

// NewForPackages returns a new instance of Updater which only updates a selected set of packages.
func NewForPackages(registry *registry.Registry, packages []string) (*Updater, error) {
	u := New(registry)
	u.packages = packages
	for _, pkgName := range u.packages {
		if _, ok := u.Registry.Packages[pkgName]; !ok {
			return nil, ErrNoSuchPackage
		}
	}
	return u, nil
}

// Update updates the registry.
func (u *Updater) Update(progress, verbose, force bool) (updated map[string]string, unchanged []string, norule []string, skipped []string, errored map[string]error) {
	updated = map[string]string{}
	unchanged = []string{}
	norule = []string{}
	skipped = []string{}
	errored = map[string]error{}
	// TODO: multithreaded for loop.
	helpers.Verbose = verbose
	i := 0
	for pkgName := range u.Registry.Packages {
		i++
		if progress {
			fmt.Printf("[%d/%d] Checking %s\n", i, len(u.Registry.Packages), pkgName)
		}

		if len(u.packages) > 0 && !includes(u.packages, pkgName) {
			skipped = append(skipped, pkgName)
			if verbose {
				fmt.Printf("  Skipped %s because not on list of packages to update\n", pkgName)
			}
			continue
		}

		v, d, ok := rules.GetRule(pkgName)
		if !ok {
			norule = append(norule, pkgName)
			if verbose {
				fmt.Printf("  No rule for %s\n", pkgName)
			}
			continue
		}

		if verbose {
			fmt.Printf("  Getting version for %s\n", pkgName)
		}
		version, err := v()
		if err != nil {
			errored[pkgName] = err
			if verbose {
				fmt.Printf("  Error checking version for %s: %v\n", pkgName, err)
			}
			continue
		}
		if verbose {
			fmt.Printf("  Version for %s: %s -> %s\n", pkgName, u.Registry.Packages[pkgName].Version, version)
		}

		if !force && u.Registry.Packages[pkgName].Version != "latest" && u.Registry.Packages[pkgName].Version == version {
			unchanged = append(unchanged, pkgName)
			if verbose {
				fmt.Printf("  Skipping %s\n", pkgName)
			}
			continue
		}

		if verbose {
			fmt.Printf("  Getting links for %s\n", pkgName)
		}
		x86dl, x86_64dl, err := d(version)
		if err != nil {
			errored[pkgName] = err
			if verbose {
				fmt.Printf("  Error getting links for %s: %v\n", pkgName, err)
			}
			continue
		}
		if verbose {
			fmt.Printf("  %s: x86: %s\n", pkgName, x86dl)
			if x86_64dl != nil {
				fmt.Printf("  %s: x86_64: %s\n", pkgName, *x86_64dl)
			} else {
				fmt.Printf("  %s: x86_64: <nil>\n", pkgName)
			}
		}

		if x86dl == "" {
			errored[pkgName] = errors.New("empty x86 download link returned")
			if verbose {
				fmt.Printf("  Error parsing links for %s: %v\n", pkgName, err)
			}
			continue
		}

		do64 := false
		if x86_64dl != nil {
			if *x86_64dl == "" {
				errored[pkgName] = errors.New("empty (but not nil) x86_64 download link returned")
				if verbose {
					fmt.Printf("  Error parsing links for %s: %v\n", pkgName, err)
				}
				continue
			}
			do64 = true
		}

		tmp := u.Registry.Packages[pkgName]
		if tmp.Version == "latest" && tmp.Installer.X86 == x86dl {
			if !(do64 && *tmp.Installer.X86_64 != *x86_64dl) {
				// Not updated a package with no version
				if verbose {
					fmt.Printf("  Version for %s is latest, and download links have not changed\n", pkgName)
				}
				unchanged = append(unchanged, pkgName)
				continue
			}
		}
		tmp.Installer.X86 = x86dl
		if do64 {
			tmp.Installer.X86_64 = x86_64dl
		} else {
			tmp.Installer.X86_64 = nil
		}
		tmp.Version = version
		u.Registry.Packages[pkgName] = tmp

		if verbose {
			fmt.Printf("  Updated %s\n", pkgName)
		}
		updated[pkgName] = version
	}
	return updated, unchanged, norule, skipped, errored
}
