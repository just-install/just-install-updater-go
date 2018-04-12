package rules

import (
	"regexp"

	. "github.com/just-install/just-install-updater-go/jiup/rules/helpers"
)

func init() {
	AddRule(
		"syncthing",
		GitHubReleaseVersionExtractor("syncthing", "syncthing", regexp.MustCompile("v(.+)")),
		GitHubReleaseDownloadExtractor("syncthing", "syncthing", regexp.MustCompile(".*windows-386.*"), regexp.MustCompile(".*windows-amd64.*")),
	)
	AddRule(
		"tortoisegit",
		RegexpVersionExtractor("https://tortoisegit.org/download/", regexp.MustCompile("TortoiseGit-([0-9.]+)")),
		HTMLDownloadExtractor("https://tortoisegit.org/download/", true, "a[href$='32bit.msi']", "a[href$='64bit.msi']", "href", "href", nil, nil),
	)
}
