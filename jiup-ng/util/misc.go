package util

import "regexp"

// TODO: URL resolving helpers

// TODO: array helpers

// Re is an alias for regexp.MustCompile.
func Re(str string) *regexp.Regexp {
	return regexp.MustCompile(str)
}

// Literal creates a regexp which matches an entire string literally.
func Literal(str string) *regexp.Regexp {
	return Re("^" + regexp.QuoteMeta(str) + "$")
}
