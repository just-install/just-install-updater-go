package v

import (
	c "github.com/just-install/just-install-updater-go/jiup/rules/common"
)

// Latest always returns the version latest.
func Latest() c.VersionExtractorFunc {
	return func() (string, error) {
		return "latest", nil
	}
}

// LatestS returns the version latest with a suffix.
func LatestS(suffix string) c.VersionExtractorFunc {
	return func() (string, error) {
		return "latest" + suffix, nil
	}
}
