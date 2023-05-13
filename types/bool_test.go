package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var bt BoolType

func TestBoolIf(t *testing.T) {

	assert.Equal(t, true, bt.If(true, true, false))
	assert.Equal(t, false, bt.If(false, true, false))
}

func TestBoolUtils(t *testing.T) {
	assert.False(t, bt.And(1 > 2, 2 > 1, 1+1 > 2, 1+1 == 2))
	assert.True(t, bt.Or(1 > 2, 2 > 1, 1+1 > 2, 1+1 == 2))
	assert.True(t, bt.Not(1 > 2))
}
