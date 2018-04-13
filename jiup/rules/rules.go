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
		"bcuninstaller",
		GitHubReleaseVersionExtractor("Klocman", "Bulk-Crap-Uninstaller", regexp.MustCompile("v(.+)")),
		GitHubReleaseDownloadExtractor("Klocman", "Bulk-Crap-Uninstaller", regexp.MustCompile(".*setup.exe"), nil),
	)
	AddRule(
		"bitpay",
		GitHubReleaseVersionExtractor("bitpay", "copay", regexp.MustCompile("v(.+)")),
		GitHubReleaseDownloadExtractor("bitpay", "copay", regexp.MustCompile("BitPay.exe"), nil),
	)
	AddRule(
		"brackets",
		GitHubReleaseVersionExtractor("adobe", "brackets", regexp.MustCompile("release-(.+)")),
		GitHubReleaseDownloadExtractor("adobe", "brackets", regexp.MustCompile("Brackets.Release.*.msi"), nil),
	)
	AddRule(
		"clementine-player",
		GitHubReleaseVersionExtractor("clementine-player", "Clementine", regexp.MustCompile("(.+)")),
		GitHubReleaseDownloadExtractor("clementine-player", "Clementine", regexp.MustCompile("ClementineSetup-.*.exe"), nil),
	)
	AddRule(
		"conemu",
		GitHubReleaseVersionExtractor("Maximus5", "ConEmu", regexp.MustCompile("v(.+)")),
		GitHubReleaseDownloadExtractor("Maximus5", "ConEmu", regexp.MustCompile("ConEmuSetup.*.exe"), nil),
	)
	AddRule(
		"dbeaver",
		GitHubReleaseVersionExtractor("dbeaver", "dbeaver", regexp.MustCompile("(.+)")),
		GitHubReleaseDownloadExtractor("dbeaver", "dbeaver", regexp.MustCompile("dbeaver-ce-.+-x86-setup.exe"), regexp.MustCompile("dbeaver-ce-.+-x86_64-setup.exe")),
	)
	AddRule(
		"syncthing",
		GitHubReleaseVersionExtractor("syncthing", "syncthing", regexp.MustCompile("v(.+)")),
		GitHubReleaseDownloadExtractor("syncthing", "syncthing", regexp.MustCompile(".*windows-386.*.zip"), regexp.MustCompile(".*windows-amd64.*.zip")),
	)
	AddRule(
		"tortoisegit",
		RegexpVersionExtractor("https://tortoisegit.org/download/", regexp.MustCompile("TortoiseGit-([0-9.]+)")),
		HTMLDownloadExtractor("https://tortoisegit.org/download/", true, "a[href$='32bit.msi']", "a[href$='64bit.msi']", "href", "href", nil, nil),
	)
}
