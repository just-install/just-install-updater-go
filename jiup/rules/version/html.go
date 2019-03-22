package v

import (
	"errors"
	"regexp"
	"strings"

	c "github.com/just-install/just-install-updater-go/jiup/rules/common"
	h "github.com/just-install/just-install-updater-go/jiup/rules/helper"
)

// HTML returns a version extractor for the first match of a css selector, an attribute (or innerText for the text), and an optional regexp on the attribute.
func HTML(url string, versionSelector, versionAttr string, versionRe *regexp.Regexp) c.VersionExtractorFunc {
	return func() (string, error) {
		doc, err := h.GetDoc(nil, url, map[string]string{}, []int{200})
		if err != nil {
			return "", err
		}

		s := doc.Find(versionSelector).First()
		if s.Length() != 1 {
			return "", errors.New("could not find match for selector")
		}

		var a string
		if versionAttr == "innerText" {
			a = strings.TrimSpace(s.Text())
		} else {
			a = strings.TrimSpace(s.AttrOr(versionAttr, ""))
		}
		if a == "" {
			return "", errors.New("specified attribute is empty")
		}

		if versionRe == nil {
			return a, nil
		}

		m := versionRe.FindStringSubmatch(a)
		if len(m) != 2 || m[1] == "" {
			return "", errors.New("could not find 2nd match group for version")
		}

		return m[1], nil
	}
}
