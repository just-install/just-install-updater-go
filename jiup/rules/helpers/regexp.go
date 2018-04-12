package helpers

import (
	"regexp"
)

// RegexpVersionExtractor returns a version extractor for a regex on a url.
func RegexpVersionExtractor(url string, versionRe *regexp.Regexp) VersionExtractorFunc {
	return func() (string, error) {
		return "", ErrExtractorNotImplemented
	}
}

// RegexpDownloadExtractor returns a version extractor for a regex on a url.
func RegexpDownloadExtractor(url string, x86FileRe, x64FileRe *regexp.Regexp) func(_ string) (string, *string, error) {
	return func(_ string) (string, *string, error) {
		return "", nil, ErrExtractorNotImplemented
	}
}
