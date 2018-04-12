package rules

import . "github.com/just-install/just-install-updater-go/jiup/rules/helpers"

var rules = map[string]struct {
	v VersionExtractorFunc
	d DownloadExtractorFunc
}{}

// AddRule registers a rule.
func AddRule(pkg string, versionExtractor VersionExtractorFunc, downloadExtractor DownloadExtractorFunc) {
	if _, ok := rules[pkg]; ok {
		panic("rule for " + pkg + " already registered")
	}
	rules[pkg] = struct {
		v VersionExtractorFunc
		d DownloadExtractorFunc
	}{versionExtractor, downloadExtractor}
}

// GetRule gets a rule if it exists.
func GetRule(pkg string) (VersionExtractorFunc, DownloadExtractorFunc, bool) {
	if rule, ok := rules[pkg]; ok {
		return rule.v, rule.d, true
	}
	return nil, nil, false
}
