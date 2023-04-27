package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolIf(t *testing.T) {
	var b BoolType
	assert.Equal(t, true, b.If(true, true, false))
	assert.Equal(t, false, b.If(false, true, false))
}
