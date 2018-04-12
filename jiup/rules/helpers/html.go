package helpers

import (
	"regexp"

	"github.com/just-install/just-install-updater-go/jiup/rules"
)

// HTMLVersionExtractor returns a version extractor for a css selector, an attribute (or innerText for the text), and an optional regexp on a url.
func HTMLVersionExtractor(url string, versionSelector, versionAttr string, versionRe *regexp.Regexp) rules.VersionExtractorFunc {
	return func() (string, error) {
		return "", rules.ErrRuleNotImplemented
	}
}

// HTMLDownloadExtractor returns a version extractor for a css selector, an attribute (or innerText for the text), and an optional regexp on a url.
func HTMLDownloadExtractor(url string, x86Selector, x86_64Selector, x86Attr, x86_64Attr string, x86FileRe, x64FileRe *regexp.Regexp) rules.DownloadExtractorFunc {
	return func() (string, *string, error) {
		return "", nil, rules.ErrRuleNotImplemented
	}
}
