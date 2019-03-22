package d

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	c "github.com/just-install/just-install-updater-go/jiup/rules/common"
	h "github.com/just-install/just-install-updater-go/jiup/rules/helper"
)

// HTML returns a download extractor for the first match of a css selector, an attribute (or innerText for the text), and an optional regexp on the url (and resolves the url).
// If a match fails, the next one (if any) is tried.
func HTML(url string, x86Selector, x64Selector, x86Attr, x64Attr string, x86FileRe, x64FileRe *regexp.Regexp) c.DownloadExtractorFunc {
	return func(_ string) (*string, *string, error) {
		doc, err := h.GetDoc(nil, url, map[string]string{}, []int{200})
		if err != nil {
			return nil, nil, err
		}

		x86dl, err := doSelector(doc, x86Selector, x86Attr, url, x86FileRe)
		if err != nil {
			return nil, nil, fmt.Errorf("could not find x86 link: %v", err)
		}

		x64dl, err := doSelector(doc, x64Selector, x64Attr, url, x64FileRe)
		if err != nil {
			return nil, nil, fmt.Errorf("could not find x64 link: %v", err)
		}

		if x86dl == nil && x64dl == nil {
			return nil, nil, errors.New("at least one of x86 and x64 must exist")
		}

		return x86dl, x64dl, nil
	}
}

// HTMLA is a shorthand version of HTML for selecting a link with a href attribute without a regex. Leave the x64 selector blank if no 64-bit version.
func HTMLA(url, x86Selector, x64Selector string) c.DownloadExtractorFunc {
	return HTML(url, x86Selector, x64Selector, "href", "href", nil, nil)
}

func doSelector(doc *goquery.Document, sel, attr, baseURL string, fileRe *regexp.Regexp) (*string, error) {
	if sel == "" {
		return nil, nil
	}
	if attr == "" {
		return nil, errors.New("attr must not be empty (use innerText for the contents)")
	}

	matches := doc.Find(sel)
	if matches.Length() < 1 {
		return nil, errors.New("could not find any matches for x86 selector")
	}

	var a string
	var err error
	matches.EachWithBreak(func(_ int, match *goquery.Selection) bool {
		if attr == "innerText" {
			a = strings.TrimSpace(match.Text())
		} else {
			a = strings.TrimSpace(match.AttrOr(attr, ""))
		}
		if a == "" {
			h, _ := match.Html()
			err = errors.New("specified attribute for x86 is empty on element: " + h)
			return true // see if any other matches don't have an issue
		}
		r, erra := h.ResolveURL(baseURL, a)
		if erra != nil {
			err = fmt.Errorf("could not resolve url: %v", erra)
			return true
		}
		a = r
		if fileRe != nil {
			m := fileRe.FindStringSubmatch(a)
			if len(m) != 2 || m[1] == "" {
				err = errors.New("could not find 2nd match group for x86 link regexp")
				return true
			}
			a = m[1]
		}
		return false
	})

	if a == "" {
		return nil, err
	}
	return &a, err
}
