package jiup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncludes(t *testing.T) {
	for _, c := range []struct {
		Arr []string
		Val string
		Res bool
	}{
		{[]string{"a", "b", "c"}, "a", true},
		{[]string{"a", "b", "c"}, "b", true},
		{[]string{"a", "b", "c"}, "c", true},
		{[]string{"a", "b", "c"}, "d", false},
	} {
		if c.Res {
			assert.True(t, includes(c.Arr, c.Val))
			continue
		}
		assert.False(t, includes(c.Arr, c.Val))
	}
}
