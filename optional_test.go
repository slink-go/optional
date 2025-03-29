package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionalGetSet(t *testing.T) {
	opt := NewFloat()
	assert.True(t, opt.isEmpty)
	assert.Equal(t, 0.0, opt.Get())
	opt.Set(1.0)
	assert.False(t, opt.isEmpty)
	assert.Equal(t, 1.0, opt.Get())
}
