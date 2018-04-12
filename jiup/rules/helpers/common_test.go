package helpers

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetURL(t *testing.T) {
	for _, c := range []struct {
		C                  *http.Client
		URL                string
		Headers            map[string]string
		AcceptableStatuses []int
		Resp               []byte
		Code               int
		OK                 bool
		HasError           bool
	}{
		{nil, "http://httpbin.org/base64/aGVsbG8gd29ybGQNCg%3D%3D", map[string]string{}, []int{200}, []byte("hello world"), 200, true, false},
		{nil, "http://httpbin.org/status/400", map[string]string{}, []int{200, 400}, []byte(""), 400, true, false},
		{nil, "http://httpbin.org/status/400", map[string]string{}, []int{200}, []byte(""), 400, false, false},
	} {
		buf, code, ok, err := GetURL(c.C, c.URL, c.Headers, c.AcceptableStatuses)
		if c.HasError {
			assert.Error(t, err)
			continue
		}
		assert.Equal(t, c.Code, code)
		assert.Equal(t, c.OK, ok)
		assert.True(t, bytes.HasPrefix(buf, c.Resp))
	}
}

func TestGetDoc(t *testing.T) {
	doc, err := GetDoc(nil, "http://httpbin.org/html", map[string]string{}, []int{200})
	assert.NoError(t, err)
	assert.Equal(t, "Herman Melville - Moby-Dick", doc.Find("h1").Text())

	doc, err = GetDoc(nil, "http://httpbin.org/404/html", map[string]string{}, []int{200, 404})
	assert.NoError(t, err)
	assert.Equal(t, "Not Found", doc.Find("h1").Text())

	doc, err = GetDoc(nil, "http://httpbin.org/404/html", map[string]string{}, []int{200})
	assert.Error(t, err)
}

func TestResolveURL(t *testing.T) {
	for _, c := range []struct {
		Base     string
		Rel      string
		Res      string
		HasError bool
	}{
		{"<|://sdf", "sdf", "", true},
		{"https://www.github.com", "sdf", "https://www.github.com/sdf", false},
		{"https://www.github.com/test", "sdf", "https://www.github.com/sdf", false},
		{"https://www.github.com/test/", "sdf", "https://www.github.com/test/sdf", false},
		{"https://www.github.com/test", "../sdf", "https://www.github.com/sdf", false},
		{"https://www.github.com/test", "../../sdf", "https://www.github.com/sdf", false},
		{"https://www.github.com/test", "/sdf", "https://www.github.com/sdf", false},
		{"https://www.github.com/test", "//www.google.ca/test", "https://www.google.ca/test", false},
		{"http://www.github.com/test", "//www.google.ca/test", "http://www.google.ca/test", false},
		{"https://www.github.com/test", "http://www.google.ca", "http://www.google.ca", false},
	} {
		res, err := ResolveURL(c.Base, c.Rel)
		if c.HasError {
			assert.Error(t, err)
			continue
		}
		assert.Equal(t, c.Res, res)
	}
}
