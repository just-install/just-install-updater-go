package util

import "testing"

func TestRe(t *testing.T) {
	Re("asd") // ensure it doesn't panic
}

func TestLiteral(t *testing.T) {
	for _, tc := range []struct {
		re, in string
		out    bool
	}{
		{"", "", true},
		{"^$", "^$", true},
		{"asd", "asd", true},
		{" ", " ", true},
		{" ", "  ", false},
		{"asd", "  asd", false},
	} {
		if Literal(tc.re).MatchString(tc.in) != tc.out {
			t.Errorf("Literal(%#v) !%t for %#v", tc.re, tc.out, tc.in)
		}
	}
}
