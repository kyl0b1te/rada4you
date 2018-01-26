package rada4you

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cli := New("1")
	assert.Equal(t, "1", cli.APIKey)
}