package v

import (
	"errors"
	"fmt"
	"regexp"

	c "github.com/just-install/just-install-updater-go/jiup/rules/common"
	h "github.com/just-install/just-install-updater-go/jiup/rules/helper"
)

// Regexp returns a version extractor for the first match of a regex.
func Regexp(url string, versionRe *regexp.Regexp) c.VersionExtractorFunc {
	return func() (string, error) {
		buf, code, ok, err := h.GetURL(nil, url, map[string]string{}, []int{200})
		if err != nil {
			return "", err
		}
		if !ok {
			return "", fmt.Errorf("unexpected response status: %d", code)
		}

		m := versionRe.FindStringSubmatch(string(buf))
		if len(m) != 2 || m[1] == "" {
			return "", errors.New("could not find 2nd match group for version regexp")
		}
		return m[1], nil
	}
}
