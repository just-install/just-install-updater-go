# just-install-updater-go
Rewrite of just-install-updater in Go.

Unit tests: `go test -v ./...`

Usage:
````
Usage: just-install-updater [options] registry [packages...]

  -d, --dry-run
        Do not actually write the changes
  --help
        Show this help text
  -q, --quiet
        Do not output progress info
  -v, --verbose
        Show more output

Arguments:
  registry is the path to just-install.json
  packages are the packages to update (default is all)
````

Usage of reachability test:
````
Usage: go run jiup/rules/reachability-test/main.go [options]

  --help
        Show this help text
  -d, --no-download
        Do not test downloadability
  -l, --download-links
        Show download links
````