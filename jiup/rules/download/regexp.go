package d

import (
	"errors"
	"fmt"
	"regexp"

	c "github.com/just-install/just-install-updater-go/jiup/rules/common"
	h "github.com/just-install/just-install-updater-go/jiup/rules/helper"
)

// Regexp returns a version extractor for the first match of a regex (and resolves the url).
func Regexp(url string, x86FileRe, x64FileRe *regexp.Regexp) c.DownloadExtractorFunc {
	return func(_ string) (*string, *string, error) {
		if x86FileRe == nil && x64FileRe == nil {
			return nil, nil, errors.New("at least one of x86 and x64 regexps must be defined")
		}

		buf, code, ok, err := h.GetURL(nil, url, map[string]string{}, []int{200})
		if err != nil {
			return nil, nil, err
		}
		if !ok {
			return nil, nil, fmt.Errorf("unexpected response status: %d", code)
		}

		var x86dl, x64dl *string
		if x86FileRe != nil {
			m := x86FileRe.FindStringSubmatch(string(buf))
			if len(m) != 2 || m[1] == "" {
				return x86dl, x64dl, errors.New("could not find 2nd match group for x86 download link regexp")
			}
			x86dls, err := h.ResolveURL(url, m[1])
			if err != nil {
				return x86dl, x64dl, err
			}
			x86dl = h.StrPtr(x86dls)
		}
		if x64FileRe != nil {
			m := x64FileRe.FindStringSubmatch(string(buf))
			if len(m) != 2 || m[1] == "" {
				return x86dl, x64dl, errors.New("could not find 2nd match group for x86 download link regexp")
			}
			x64dls, err := h.ResolveURL(url, m[1])
			if err != nil {
				return x86dl, x64dl, err
			}
			x64dl = h.StrPtr(x64dls)
		}

		return x86dl, x64dl, nil
	}
}
