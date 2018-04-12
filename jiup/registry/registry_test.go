package registry

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/just-install.json")
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	r, err := NewFromJSON(buf)
	assert.NoError(t, err)
	assert.NotNil(t, r)

	bufn, err := r.GetJSON()
	assert.NoError(t, err)
	assert.NotNil(t, bufn)

	cmdA := exec.Command("jq", "-S", ".")
	cmdA.Stdin = bytes.NewReader(buf)
	outA, err := cmdA.Output()
	assert.NoError(t, err)

	cmdB := exec.Command("jq", "-S", ".")
	cmdB.Stdin = bytes.NewReader(bufn)
	outB, err := cmdB.Output()
	assert.NoError(t, err)

	assert.Equal(t, string(outA), string(outB), "normalized JSON should be the same")
}
