package w

import (
	"strings"

	c "github.com/just-install/just-install-updater-go/jiup/rules/common"
	h "github.com/just-install/just-install-updater-go/jiup/rules/helper"
)

// UnderscoreToDot wraps a version extractor and replaces underscores with dots.
func UnderscoreToDot(f c.VersionExtractorFunc) c.VersionExtractorFunc {
	return func() (string, error) {
		version, err := f()
		if err != nil {
			return "", err
		}
		return strings.Replace(version, "_", ".", -1), nil
	}
}

// AppendToURL wraps a download extractor and appends a string to each URL.
func AppendToURL(str string, f c.DownloadExtractorFunc) c.DownloadExtractorFunc {
	return func(version string) (*string, *string, error) {
		x86, x64, err := f(version)
		if err != nil {
			return nil, nil, err
		}
		if x64 != nil {
			t := *x64 + str
			x64 = &t
		}
		x86 = h.StrPtr(*x86 + str)
		return x86, x64, nil
	}
}

// SplitDownload runs a different helper for each architecture.
func SplitDownload(x86, x64 c.DownloadExtractorFunc) c.DownloadExtractorFunc {
	return func(version string) (*string, *string, error) {
		dx86, _, err := x86(version)
		if err != nil {
			return nil, nil, err
		}
		_, dx64, err := x64(version)
		if err != nil {
			return nil, nil, err
		}
		return dx86, dx64, nil
	}
}
