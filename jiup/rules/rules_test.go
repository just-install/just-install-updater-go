package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidity(t *testing.T) {
	for p, r := range rules {
		assert.NotEmpty(t, p, "rule package name should not be empty")
		assert.NotNil(t, r.v, "version extractor should not be nil")
		assert.NotNil(t, r.d, "download extractor should not be nil")
	}
}
