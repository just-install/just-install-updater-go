package helpers

import (
	"regexp"

	"github.com/just-install/just-install-updater-go/jiup/rules"
)

// RegexpVersionExtractor returns a version extractor for a regex on a url.
func RegexpVersionExtractor(url string, versionRe *regexp.Regexp) rules.VersionExtractorFunc {
	return func() (string, error) {
		return "", rules.ErrRuleNotImplemented
	}
}

// RegexpDownloadExtractor returns a version extractor for a regex on a url.
func RegexpDownloadExtractor(url string, x86FileRe, x64FileRe *regexp.Regexp) rules.DownloadExtractorFunc {
	return func() (string, *string, error) {
		return "", nil, rules.ErrRuleNotImplemented
	}
}
