package d

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/just-install/just-install-updater-go/jiup/rules/c"
	"github.com/just-install/just-install-updater-go/jiup/rules/h"
)

// GitHubRelease returns a version extractor for a GitHub release. x64Re can be nil.
func GitHubRelease(repo string, x86FileRe, x64FileRe *regexp.Regexp) c.DownloadExtractorFunc {
	return func(_ string) (*string, *string, error) {
		if x86FileRe == nil && x64FileRe == nil {
			return nil, nil, errors.New("at least one of x86 and x64 regexps must be defined")
		}

		// scrape to avoid limit
		doc, err := h.GetDoc(nil, fmt.Sprintf("https://github.com/%s/releases/latest", repo), map[string]string{}, []int{200})
		if err != nil {
			return nil, nil, err
		}

		files := [][]string{}
		err = nil
		doc.Find(".release").First().Find(".Details-element:contains('Assets') .Box a[href][href*='download']").EachWithBreak(func(_ int, s *goquery.Selection) bool {
			href := strings.TrimSpace(s.AttrOr("href", ""))
			if href == "" {
				err = errors.New("could not extract href from release asset")
				return false
			}
			href, err = h.ResolveURL(fmt.Sprintf("https://github.com/%s/releases/latest", repo), href)
			if err != nil {
				return false
			}
			spl := strings.Split(href, "/")
			fname := spl[len(spl)-1]
			if fname == "" {
				err = errors.New("could not extract filename from release asset")
				return false
			}
			if strings.HasSuffix(fname, ".sig") {
				// Skip signature files
				return true
			}
			if strings.HasSuffix(fname, ".sha1") || strings.HasSuffix(fname, ".sha256") || strings.HasSuffix(fname, ".md5") {
				// Skip sha files
				return true
			}
			files = append(files, []string{href, fname})
			return true
		})
		if err != nil {
			return nil, nil, err
		}
		if len(files) == 0 {
			return nil, nil, errors.New("could not extract list of assets")
		}

		var x86dl, x64dl *string
		if x86FileRe != nil {
			for _, file := range files {
				if x86FileRe.MatchString(file[1]) {
					x86dl = h.StrPtr(file[0])
					break
				}
			}
			if x86dl == nil {
				return nil, nil, errors.New("could not find asset filename match for x86")
			}
		}

		if x64FileRe != nil {
			for _, file := range files {
				if x64FileRe.MatchString(file[1]) {
					x64dl = h.StrPtr(file[0])
					break
				}
			}
			if x64dl == nil {
				return nil, nil, errors.New("could not find asset filename match for x64")
			}
		}

		if x86dl == nil && x64dl == nil {
			return nil, nil, errors.New("could not find match for x86 or x64")
		}
		return x86dl, x64dl, nil
	}
}
