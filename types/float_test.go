package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatIf(t *testing.T) {
	var f FloatType
	assert.Equal(t, 1.11, f.If(true, 1.11, 0.00))
	assert.Equal(t, 0.00, f.If(false, 1.11, 0.00))
}
