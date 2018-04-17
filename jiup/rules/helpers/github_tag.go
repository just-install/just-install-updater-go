package helpers

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// GitHubTagVersionExtractor returns a version extractor for a GitHub tag.
func GitHubTagVersionExtractor(username, repo string, tagRe *regexp.Regexp) VersionExtractorFunc {
	return func() (string, error) {
		if tagRe == nil {
			return "", errors.New("tag regex is nil")
		}

		// scrape to avoid limit
		doc, err := GetDoc(nil, fmt.Sprintf("https://github.com/%s/%s/tags", username, repo), map[string]string{}, []int{200})
		if err != nil {
			return "", err
		}

		tag := strings.TrimSpace(doc.Find(".releases-tag-list .tag-info .tag-name").First().Text())
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
