package helpers

import (
	"errors"
	"fmt"
	"regexp"
)

// RegexpVersionExtractor returns a version extractor for a regex.
func RegexpVersionExtractor(url string, versionRe *regexp.Regexp) VersionExtractorFunc {
	return func() (string, error) {
		buf, code, ok, err := GetURL(nil, url, map[string]string{}, []int{200})
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

// RegexpDownloadExtractor returns a version extractor for a regex (and resolves the url).
func RegexpDownloadExtractor(url string, x86FileRe, x64FileRe *regexp.Regexp) func(_ string) (string, *string, error) {
	return func(_ string) (string, *string, error) {
		buf, code, ok, err := GetURL(nil, url, map[string]string{}, []int{200})
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

		x86dl, err = ResolveURL(url, x86dl)
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

		x64dl, err = ResolveURL(url, x64dl)
		if err != nil {
			return "", nil, err
		}

		return x86dl, &x64dl, nil
	}
}
