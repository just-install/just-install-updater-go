package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidity(t *testing.T) {
	for p, r := range rules {
		assert.NotEmpty(t, p, "rule package name should not be empty")
		assert.NotNil(t, r.V, "version extractor should not be nil")
		assert.NotNil(t, r.D, "download extractor should not be nil")
	}
}
