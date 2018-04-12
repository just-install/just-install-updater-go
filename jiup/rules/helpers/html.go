package helpers

import (
	"regexp"
)

// HTMLVersionExtractor returns a version extractor for a css selector, an attribute (or innerText for the text), and an optional regexp on a url.
func HTMLVersionExtractor(url string, versionSelector, versionAttr string, versionRe *regexp.Regexp) VersionExtractorFunc {
	return func() (string, error) {
		return "", ErrExtractorNotImplemented
	}
}

// HTMLDownloadExtractor returns a version extractor for a css selector, an attribute (or innerText for the text), and an optional regexp on a url.
func HTMLDownloadExtractor(url string, x86Selector, x86_64Selector, x86Attr, x86_64Attr string, x86FileRe, x64FileRe *regexp.Regexp) DownloadExtractorFunc {
	return func(_ string) (string, *string, error) {
		return "", nil, ErrExtractorNotImplemented
	}
}
