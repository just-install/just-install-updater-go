package d

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/just-install/just-install-updater-go/jiup/rules/c"
	"github.com/just-install/just-install-updater-go/jiup/rules/h"
)

// Regexp returns a version extractor for the first match of a regex (and resolves the url).
func Regexp(url string, x86FileRe, x64FileRe *regexp.Regexp) c.DownloadExtractorFunc {
	return func(_ string) (string, *string, error) {
		buf, code, ok, err := h.GetURL(nil, url, map[string]string{}, []int{200})
		if err != nil {
			return "", nil, err
		}
		if !ok {
			return "", nil, fmt.Errorf("unexpected response status: %d", code)
		}

		m := x86FileRe.FindStringSubmatch(string(buf))
		if len(m) != 2 || m[1] == "" {
			return "", nil, errors.New("could not find 2nd match group for x86 download link regexp")
		}
		x86dl := m[1]

		x86dl, err = h.ResolveURL(url, x86dl)
		if err != nil {
			return "", nil, err
		}

		if x64FileRe == nil {
			return x86dl, nil, nil
		}

		m = x64FileRe.FindStringSubmatch(string(buf))
		if len(m) != 2 || m[1] == "" {
			return "", nil, errors.New("could not find 2nd match group for x64 download link regexp")
		}
		x64dl := m[1]

		x64dl, err = h.ResolveURL(url, x64dl)
		if err != nil {
			return "", nil, err
		}

		return x86dl, &x64dl, nil
	}
}
