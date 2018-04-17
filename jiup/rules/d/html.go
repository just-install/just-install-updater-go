package d

import (
	"errors"
	"regexp"
	"strings"

	"github.com/just-install/just-install-updater-go/jiup/rules/c"
	"github.com/just-install/just-install-updater-go/jiup/rules/h"
)

// HTML returns a download extractor for the first match of a css selector, an attribute (or innerText for the text), and an optional regexp on the url (and resolves the url).
func HTML(url string, hasx86_64 bool, x86Selector, x86_64Selector, x86Attr, x86_64Attr string, x86FileRe, x64FileRe *regexp.Regexp) c.DownloadExtractorFunc {
	if !hasx86_64 && x86_64Selector != "" {
		panic("x86_64Selector defined while hasx86_64 is false")
	} else if !hasx86_64 && x86_64Attr != "" {
		panic("x86_64Attr defined while hasx86_64 is false")
	} else if hasx86_64 && x86_64Selector == "" {
		panic("x86_64Selector empty while hasx86_64 is true")
	} else if hasx86_64 && x86_64Attr == "" {
		panic("x86_64Attr empty while hasx86_64 is true")
	}

	return func(_ string) (string, *string, error) {
		doc, err := h.GetDoc(nil, url, map[string]string{}, []int{200})
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

		a, err = h.ResolveURL(url, a)
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

		a, err = h.ResolveURL(url, a)
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

// HTMLA is a shorthand version of HTML for selecting a link with a href attribute without a regex. Leave the x64 selector blank if no 64-bit version.
func HTMLA(url, x86Selector, x86_64Selector string) c.DownloadExtractorFunc {
	if x86_64Selector == "" {
		return HTML(url, false, x86Selector, "", "href", "", nil, nil)
	}
	return HTML(url, true, x86Selector, x86_64Selector, "href", "href", nil, nil)
}
