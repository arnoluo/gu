package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var at AnyType

func TestAnyIf(t *testing.T) {
	assert.Equal(t, 1, at.If(true, 1, 0))
	assert.Equal(t, 0, at.If(false, 1, 0))
	assert.Equal(t, "trueValue", at.If(true, "trueValue", "falseValue"))
	assert.Equal(t, "falseValue", at.If(false, "trueValue", "falseValue"))
}

// response type
type Rsp struct {
	Code int `json:"code"`
}

// response type with ts(timestamp)
type Rspt struct {
	Code int   `json:"code"`
	Ts   int64 `json:"ts"`
}

func TestStructTo(t *testing.T) {
	jsonRspt := Rspt{
		Code: 1,
		Ts:   123,
	}

	var jsonRsp Rsp
	at.StructTo(jsonRspt, &jsonRsp)
	assert.Equal(t, jsonRsp, Rsp{
		Code: 1,
	})

	var jRspt Rspt
	at.StructTo(jsonRsp, &jRspt)
	assert.Equal(t, jRspt, Rspt{
		Code: 1,
		Ts:   0,
	})
}

func TestInArray(t *testing.T) {
	assert.True(t, at.InArray(1, []int{1, 2, 3}))
	var a float64 = 1
	assert.True(t, at.InArray(a, []float64{1, 2}))
	// not recommanded
	assert.False(t, at.InArray(1, []float64{1, 2}))
	assert.True(t, at.InArray("a", []string{"b", "a"}))
	assert.True(t, at.InArray("1", []string{"1", "2"}))
	assert.False(t, at.InArray("1", []int{1, 2}))
	var b any = 1
	assert.True(t, at.InArray(b, []int{1, 2, 3}))
}
