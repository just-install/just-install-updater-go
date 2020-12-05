# just-install-updater-go
Rewrite of just-install-updater in Go.

Unit tests: `go test -v ./...`

Usage:

```
Usage: just-install-updater [options] registry [packages...]

  -c, --commit-message-file string   If set, jiup-go will save a commit message describing the changes to a file.
  -d, --dry-run                      Do not actually write the changes
  -f, --force                        Update all entries including ones with a matching version
      --help                         Show this help text
  -q, --quiet                        Do not output progress info
  -b, --read-broken string           If set, jiup-go will ignore rules listed in the specified file
  -v, --verbose                      Show more output

Arguments:
  registry is the path to just-install.json
  packages are the packages to update (default is all)
```

Usage of reachability test:

```
Usage: reachability-test [options] [packages...]

  -l, --download-links        Show download links
      --help                  Show this help text
  -d, --no-download           Do not test downloadability
  -b, --write-broken string   If set, broken rules will be written into the specified file
```