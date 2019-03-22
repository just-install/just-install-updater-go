package h

import "regexp"

// Re is an alias for regexp.MustCompile.
func Re(str string) *regexp.Regexp {
	return regexp.MustCompile(str)
}
