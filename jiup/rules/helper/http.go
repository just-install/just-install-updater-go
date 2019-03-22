package h

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
