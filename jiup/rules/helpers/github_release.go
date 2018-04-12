package helpers

import (
	"regexp"
)

// GitHubReleaseVersionExtractor returns a version extractor for a GitHub release.
func GitHubReleaseVersionExtractor(username, repo string, tagRe *regexp.Regexp) VersionExtractorFunc {
	return func() (string, error) {
		return "", ErrExtractorNotImplemented
	}
}

// GitHubReleaseDownloadExtractor returns a version extractor for a GitHub release. x64Re can be nil.
func GitHubReleaseDownloadExtractor(username, repo string, x86FileRe, x64FileRe *regexp.Regexp) DownloadExtractorFunc {
	return func(_ string) (string, *string, error) {
		return "", nil, ErrExtractorNotImplemented
	}
}
