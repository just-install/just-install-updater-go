package v

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	c "github.com/just-install/just-install-updater-go/jiup/rules/common"
	h "github.com/just-install/just-install-updater-go/jiup/rules/helper"
)

// GitHubTag returns a version extractor for a GitHub tag.
func GitHubTag(repo string, tagRe *regexp.Regexp) c.VersionExtractorFunc {
	return func() (string, error) {
		if tagRe == nil {
			return "", errors.New("tag regex is nil")
		}

		// scrape to avoid limit
		doc, err := h.GetDoc(nil, fmt.Sprintf("https://github.com/%s/tags", repo), map[string]string{}, []int{200})
		if err != nil {
			return "", err
		}

		tag := strings.TrimSpace(doc.Find(".commit.Details .commit-title a").First().Text())
		if tag == "" {
			return "", errors.New("could not find tag from GitHub")
		}

		m := tagRe.FindStringSubmatch(tag)
		if len(m) != 2 || m[1] == "" {
			return "", errors.New("could not find 2nd match group for tag regexp")
		}

		return m[1], nil
	}
}

// GitHubRelease returns a version extractor for a GitHub release.
func GitHubRelease(repo string, tagRe *regexp.Regexp) c.VersionExtractorFunc {
	return func() (string, error) {
		if tagRe == nil {
			return "", errors.New("tag regex is nil")
		}

		// scrape to avoid limit
		doc, err := h.GetDoc(nil, fmt.Sprintf("https://github.com/%s/releases/latest", repo), map[string]string{}, []int{200})
		if err != nil {
			return "", err
		}

		tag := strings.TrimSpace(doc.Find(".release .octicon-tag+span").First().Text())
		if tag == "" {
			return "", errors.New("could not find tag from GitHub")
		}

		m := tagRe.FindStringSubmatch(tag)
		if len(m) != 2 || m[1] == "" {
			return "", errors.New("could not find 2nd match group for tag regexp")
		}

		return m[1], nil
	}
}
