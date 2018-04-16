package helpers

import (
	"errors"
	"regexp"
	"strings"
)

// HTMLVersionExtractor returns a version extractor for a css selector, an attribute (or innerText for the text), and an optional regexp on the attribute.
func HTMLVersionExtractor(url string, versionSelector, versionAttr string, versionRe *regexp.Regexp) VersionExtractorFunc {
	return func() (string, error) {
		doc, err := GetDoc(nil, url, map[string]string{}, []int{200})
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

// HTMLDownloadExtractor returns a download extractor for a css selector, an attribute (or innerText for the text), and an optional regexp on the url (and resolves the url).
func HTMLDownloadExtractor(url string, hasx86_64 bool, x86Selector, x86_64Selector, x86Attr, x86_64Attr string, x86FileRe, x64FileRe *regexp.Regexp) DownloadExtractorFunc {
	if !hasx86_64 && x86_64Selector != "" {
		panic("x86_64 selector defined while hasx86_64 is false")
	}
	return func(_ string) (string, *string, error) {
		doc, err := GetDoc(nil, url, map[string]string{}, []int{200})
		if err != nil {
			return "", nil, err
		}

		var x86dl, x64dl string

		s := doc.Find(x86Selector).First()
		if s.Length() != 1 {
			return "", nil, errors.New("could not find match for x86 selector")
		}

		var a string
		if x86Attr == "innerText" {
			a = strings.TrimSpace(s.Text())
		} else {
			a = strings.TrimSpace(s.AttrOr(x86Attr, ""))
		}
		if a == "" {
			return "", nil, errors.New("specified attribute for x86 is empty")
		}

		a, err = ResolveURL(url, a)
		if err != nil {
			return "", nil, err
		}

		if x86FileRe != nil {
			m := x86FileRe.FindStringSubmatch(a)
			if len(m) != 2 || m[1] == "" {
				return "", nil, errors.New("could not find 2nd match group for x86 link regexp")
			}
			a = m[1]
		}

		x86dl = a

		if !hasx86_64 {
			return x86dl, nil, nil
		}

		s = doc.Find(x86_64Selector).First()
		if s.Length() != 1 {
			return "", nil, errors.New("could not find match for x86_64 selector")
		}

		a = ""
		if x86_64Attr == "innerText" {
			a = strings.TrimSpace(s.Text())
		} else {
			a = strings.TrimSpace(s.AttrOr(x86_64Attr, ""))
		}
		if a == "" {
			return "", nil, errors.New("specified attribute for x86_64 is empty")
		}

		a, err = ResolveURL(url, a)
		if err != nil {
			return "", nil, err
		}

		if x64FileRe != nil {
			m := x64FileRe.FindStringSubmatch(a)
			if len(m) != 2 || m[1] == "" {
				return "", nil, errors.New("could not find 2nd match group for x86_64 link regexp")
			}
			a = m[1]
		}

		x64dl = a

		return x86dl, &x64dl, nil
	}
}
