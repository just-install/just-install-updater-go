package helpers

import (
	"regexp"

	"github.com/just-install/just-install-updater-go/jiup/rules"
)

// GitHubReleaseVersionExtractor returns a version extractor for a GitHub release.
func GitHubReleaseVersionExtractor(username, repo string, tagRe *regexp.Regexp) rules.VersionExtractorFunc {
	return func() (string, error) {
		return "", rules.ErrRuleNotImplemented
	}
}

// GitHubReleaseDownloadExtractor returns a version extractor for a GitHub release. x64Re can be nil.
func GitHubReleaseDownloadExtractor(username, repo string, x86FileRe, x64FileRe *regexp.Regexp) rules.DownloadExtractorFunc {
	return func(_ string) (string, *string, error) {
		return "", nil, rules.ErrRuleNotImplemented
	}
}
