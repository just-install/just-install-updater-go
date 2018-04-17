package helpers

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type resps struct {
	Buf  []byte
	Code int
	OK   bool
	Err  error
}
type reqi string

var cache = map[reqi]resps{}

func mkreqi(url string, headers map[string]string, acceptableStatuses []int) reqi {
	return reqi(fmt.Sprintf("%#v;;;%#v;;;%#v", url, headers, acceptableStatuses))
}

// GetURL gets a url. The client is optional.
func GetURL(c *http.Client, url string, headers map[string]string, acceptableStatuses []int) ([]byte, int, bool, error) {
	ri := mkreqi(url, headers, acceptableStatuses)
	if r, ok := cache[ri]; ok {
		return r.Buf, r.Code, r.OK, r.Err
	}

	if c == nil {
		c = &http.Client{
			Timeout: time.Duration(time.Second * 10),
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		cache[ri] = resps{nil, 0, false, err}
		return nil, 0, false, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		cache[ri] = resps{nil, 0, false, err}
		return nil, 0, false, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cache[ri] = resps{nil, 0, false, err}
		return nil, 0, false, err
	}

	a := false
	for _, s := range acceptableStatuses {
		if s == resp.StatusCode {
			a = true
			break
		}
	}

	cache[ri] = resps{buf, resp.StatusCode, a, nil}
	return buf, resp.StatusCode, a, nil
}

// GetDoc gets a goquery doc from a url.
func GetDoc(c *http.Client, url string, headers map[string]string, acceptableStatuses []int) (*goquery.Document, error) {
	buf, s, a, err := GetURL(c, url, headers, acceptableStatuses)
	if err != nil {
		return nil, err
	}
	if !a {
		return nil, fmt.Errorf("unexpected response status: %d", s)
	}
	return goquery.NewDocumentFromReader(bytes.NewReader(buf))
}

// ResolveURL resolves a relative url.
func ResolveURL(base, rel string) (string, error) {
	urlP, err := url.Parse(rel)
	if err != nil {
		return "", err
	}

	baseP, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	resolved := baseP.ResolveReference(urlP)

	return resolved.String(), nil
}

// Re is an alias for regexp.MustCompile.
func Re(str string) *regexp.Regexp {
	return regexp.MustCompile(str)
}

// UnderscoreToDot wraps a version extractor and replaces underscores with dots.
func UnderscoreToDot(f VersionExtractorFunc) VersionExtractorFunc {
	return func() (string, error) {
		version, err := f()
		if err != nil {
			return "", err
		}
		return strings.Replace(version, "_", ".", -1), nil
	}
}

// AppendToURL wraps a download extractor and appends a string to each URL.
func AppendToURL(str string, f DownloadExtractorFunc) DownloadExtractorFunc {
	return func(version string) (string, *string, error) {
		x86, x64, err := f(version)
		if err != nil {
			return "", nil, err
		}
		if x64 != nil {
			t := *x64 + str
			x64 = &t
		}
		x86 = x86 + str
		return x86, x64, nil
	}
}

// DisableHTTPSCheck disables https certificate checking for the default client.
func DisableHTTPSCheck() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

// EnableHTTPSCheck reenables https certificate checking for the default client.
func EnableHTTPSCheck() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
}

// NoHTTPSForVersionExtractor wraps a VersionExtractorFunc to disable HTTPS checking.
func NoHTTPSForVersionExtractor(f VersionExtractorFunc) VersionExtractorFunc {
	return func() (string, error) {
		DisableHTTPSCheck()
		version, err := f()
		EnableHTTPSCheck()
		return version, err
	}
}

// NoHTTPSForDownloadExtractor wraps a DownloadExtractorFunc to disable HTTPS checking.
func NoHTTPSForDownloadExtractor(f DownloadExtractorFunc) DownloadExtractorFunc {
	return func(version string) (string, *string, error) {
		DisableHTTPSCheck()
		x86, x64, err := f(version)
		EnableHTTPSCheck()
		return x86, x64, err
	}
}
