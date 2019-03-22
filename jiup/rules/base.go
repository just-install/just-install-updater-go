package rules

import (
	"errors"
	"fmt"
	"strings"

	c "github.com/just-install/just-install-updater-go/jiup/rules/common"
)

// R represents a rule for an application.
type R struct {
	V c.VersionExtractorFunc
	D c.DownloadExtractorFunc
}

var rules = map[string]R{}

// Rule registers a rule.
func Rule(pkg string, versionExtractor c.VersionExtractorFunc, downloadExtractor c.DownloadExtractorFunc) {
	if _, ok := rules[pkg]; ok {
		panic("rule for " + pkg + " already registered")
	}
	rules[pkg] = R{wrapV(versionExtractor), wrapD(downloadExtractor)}
}

// GetRule gets a rule if it exists.
func GetRule(pkg string) (c.VersionExtractorFunc, c.DownloadExtractorFunc, bool) {
	if rule, ok := rules[pkg]; ok {
		return rule.V, rule.D, true
	}
	return nil, nil, false
}

// GetRules gets all rules.
func GetRules() map[string]R {
	return rules
}

func wrapV(f c.VersionExtractorFunc) c.VersionExtractorFunc {
	return func() (version string, err error) {
		version, err = f()
		if err != nil {
			return "", err
		}
		if strings.TrimSpace(version) == "" {
			return "", errors.New("version is empty")
		}
		return version, nil
	}
}

func wrapD(f c.DownloadExtractorFunc) c.DownloadExtractorFunc {
	return func(version string) (x86 *string, x86_64 *string, err error) {
		x86, x86_64, err = f(version)
		if x86 != nil {
			if err != nil {
				return nil, nil, err
			}
			if strings.TrimSpace(*x86) == "" {
				return nil, nil, errors.New("x86 link is empty")
			}
			if !strings.HasPrefix(*x86, "http") {
				return nil, nil, fmt.Errorf("x86 link (%s) does not start with http", *x86)
			}
		}
		if x86_64 != nil {
			if strings.TrimSpace(*x86_64) == "" {
				return nil, nil, errors.New("x86_64 link is empty")
			}
			if !strings.HasPrefix(*x86_64, "http") {
				return nil, nil, fmt.Errorf("x86_64 link (%s) does not start with http", *x86_64)
			}
		}
		return x86, x86_64, nil
	}
}
