package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tcase struct {
	URL      string
	Code     int
	Mime     string
	HasError bool
}

func TestTestDL(t *testing.T) {
	for _, c := range []tcase{
		{"https://httpbin.org/status/200", 200, "text/html", false},
		{"https://httpbin.org/status/404", 404, "text/html", false},
		{"http://127.0.0.1:3453", 0, "", true},
		{"https://httpbin.org/get", 200, "application/json", false},
	} {
		code, mime, err := testDL(c.URL)
		if c.HasError {
			assert.Error(t, err)
			continue
		}
		assert.NoError(t, err)
		assert.Equal(t, c.Code, code)
		assert.Contains(t, mime, c.Mime)
	}
}
