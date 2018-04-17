package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListify(t *testing.T) {
	for o, i := range map[string][]string{
		"":               {},
		"a":              {"a"},
		"a and b":        {"a", "b"},
		"a, b, and c":    {"a", "b", "c"},
		"a, b, c, and d": {"a", "b", "c", "d"},
	} {
		assert.Equal(t, o, listify(i))
	}
}
