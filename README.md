# just-install-updater-go
POC of just-install-updater written in Go.

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
  packages are the packages to update (default is all)````