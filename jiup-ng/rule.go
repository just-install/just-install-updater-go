package jiup

import "regexp"

// Rule contains a Versioner and a Downloader for a specific package. Note that
// rules must support being called multiple times and possibly concurrently (i.e.
// the rule itself should be read-only, and all data should be stored in RuleData).
type Rule interface {
	Versioner
	Downloader
}

// RuleMix combines a separate Versioner and a Downloader into a Rule. It is
// using a function instead of a composite struct to prevent warnings all over
// the rules about "composite literal uses unkeyed fields".
func RuleMix(v Versioner, d Downloader) Rule {
	return struct {
		Versioner
		Downloader
	}{v, d}
}

// Versioner gets the version info.
type Versioner interface {
	// Version gets the latest version for a package.
	Version(data *RuleData) (string, error)
}

// Downloader gets download links.
type Downloader interface {
	// Download gets the download links for a package.
	Download(version string, data *RuleData) (LinkMap, error)
}

// LinkMap contains links for different architectures.
type LinkMap map[Architecture]string

// RegexpMap contains regular expressions for each Architecture and is intended
// for use in sources.
type RegexpMap map[Architecture]*regexp.Regexp

// Architecture is an enum of the supported architectures and should be used in
// all helpers wherever possible (rather than hardcoding arguments/return values
// for specific ones).
type Architecture string

// Supported architectures.
const (
	Arch32 Architecture = "x86"
	Arch64 Architecture = "x86_64"
)
