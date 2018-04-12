package jiup

import (
	"errors"

	"github.com/just-install/just-install-updater-go/jiup/rules"

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
func (u *Updater) Update(progress, verbose bool) (updated map[string]string, unchanged []string, norule []string, skipped []string, errored map[string]error) {
	updated = map[string]string{}
	unchanged = []string{}
	norule = []string{}
	skipped = []string{}
	errored = map[string]error{}
	// TODO: multithreaded for loop.
	for pkgName := range u.Registry.Packages {
		if len(u.packages) > 0 && !includes(u.packages, pkgName) {
			skipped = append(skipped, pkgName)
			continue
		}

		v, d, ok := rules.GetRule(pkgName)
		if !ok {
			norule = append(norule, pkgName)
			continue
		}

		version, err := v()
		if err != nil {
			errored[pkgName] = err
			continue
		}

		if u.Registry.Packages[pkgName].Version != "latest" && u.Registry.Packages[pkgName].Version == version {
			unchanged = append(unchanged, pkgName)
			continue
		}

		x86dl, x86_64dl, err := d(version)
		if err != nil {
			errored[pkgName] = err
			continue
		}

		if x86dl == "" {
			errored[pkgName] = errors.New("empty x86 download link returned")
			continue
		}

		do64 := false
		if x86_64dl != nil {
			if *x86_64dl == "" {
				errored[pkgName] = errors.New("empty (but not nil) x86_64 download link returned")
				continue
			}
			do64 = true
		}

		tmp := u.Registry.Packages[pkgName]
		if tmp.Version == "latest" && tmp.Installer.X86 == x86dl {
			if !(do64 && *tmp.Installer.X86_64 != *x86_64dl) {
				// Not updated a package with no version
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

		updated[pkgName] = version
	}
	return updated, unchanged, norule, skipped, errored
}
