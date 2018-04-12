package rules

import (
	"regexp"
	"strings"

	. "github.com/just-install/just-install-updater-go/jiup/rules/helpers"
)

func init() {
	AddRule(
		"7zip",
		RegexpVersionExtractor("https://7-zip.org/download.html", regexp.MustCompile("Download 7-Zip ([0-9][0-9].[0-9][0-9])")),
		func(version string) (string, *string, error) {
			x86dl := "https://www.7-zip.org/a/7z" + strings.Replace(version, ".", "", -1) + ".msi"
			x64dl := "https://www.7-zip.org/a/7z" + strings.Replace(version, ".", "", -1) + "-x64.msi"
			return x86dl, &x64dl, nil
		},
	)
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
