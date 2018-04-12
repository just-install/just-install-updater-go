package rules

// VersionExtractorFunc represents a function which extracts the version.
type VersionExtractorFunc func() (version string, err error)

// DownloadExtractorFunc represents a function which extracts a download link for a version.
// It can return an nil string pointer for x86_64 if not available.
type DownloadExtractorFunc func() (x86 string, x86_64 *string, err error)

var rules map[string]struct {
	v VersionExtractorFunc
	d DownloadExtractorFunc
}

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
