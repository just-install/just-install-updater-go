package w

import (
	"crypto/tls"
	"net/http"

	"github.com/just-install/just-install-updater-go/jiup/rules/c"
)

// DisableHTTPSCheck disables https certificate checking for the default client.
func DisableHTTPSCheck() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

// EnableHTTPSCheck reenables https certificate checking for the default client.
func EnableHTTPSCheck() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
}

// NoHTTPSForVersionExtractor wraps a VersionExtractorFunc to disable HTTPS checking.
func NoHTTPSForVersionExtractor(f c.VersionExtractorFunc) c.VersionExtractorFunc {
	return func() (string, error) {
		DisableHTTPSCheck()
		version, err := f()
		EnableHTTPSCheck()
		return version, err
	}
}

// NoHTTPSForDownloadExtractor wraps a DownloadExtractorFunc to disable HTTPS checking.
func NoHTTPSForDownloadExtractor(f c.DownloadExtractorFunc) c.DownloadExtractorFunc {
	return func(version string) (string, *string, error) {
		DisableHTTPSCheck()
		x86, x64, err := f(version)
		EnableHTTPSCheck()
		return x86, x64, err
	}
}
