package gu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	assert.Equal(t, "trueValue", St.If(true, "trueValue", "falseValue"))
	assert.Equal(t, "falseValue", St.If(false, "trueValue", "falseValue"))
	assert.Equal(t, int64(1), It.If(true, 1, 0))
	assert.Equal(t, int64(0), It.If(false, 1, 0))
	assert.Equal(t, true, Bt.If(true, true, false))
	assert.Equal(t, false, Bt.If(false, true, false))
	assert.Equal(t, 1.11, Ft.If(true, 1.11, 0.00))
	assert.Equal(t, 0.00, Ft.If(false, 1.11, 0.00))
	assert.Equal(t, 1, At.If(true, 1, 0))
	assert.Equal(t, 0, At.If(false, 1, 0))
}

func TestInArray(t *testing.T) {
	strs := []string{"foo", "bar", "baz", "hello", "world", "golang", "google", "facebook", "amazon", "microsoft"}
	ints := []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	floats := []float64{1.23, 4.56, 7.89, 0.12, 3.45, 6.78}
	interfaces := []any{1, "hello", 3.1415926, true, "world"}

	assert.True(t, St.InArray("golang", strs))
	assert.True(t, It.InArray(16, ints))
	assert.True(t, Ft.InArray(6.78, floats))
	assert.True(t, At.InArray(3.1415926, interfaces))
	assert.False(t, At.InArray(2.71828, interfaces))
}
