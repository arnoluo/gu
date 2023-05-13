package types

import (
	"errors"
	"fmt"
	"math"
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

func TestAnyInt64(t *testing.T) {
	var toVal int64
	var i int8 = 123
	targetI := int64(i)
	at.Int64(i, &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(int16(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(int32(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(int64(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(int(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(uint(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(uintptr(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(uint8(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(uint16(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(uint32(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(uint64(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(float32(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(float64(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(fmt.Sprintf("%d", i), &toVal)
	assert.Equal(t, targetI, toVal)

	err := at.Int64(uint64(math.MaxInt64)+1, &toVal)
	assert.Equal(t, errors.New("gu.At.Int64() Error: out of range int64"), err)

	var a any
	err = at.Int64(a, &toVal)
	assert.Equal(t, errors.New("gu.At.Int64() Error: unsupported type"), err)
}

func TestAnyInt(t *testing.T) {
	var toVal int64
	var i int8 = 123
	targetI := int64(i)
	at.Int64(i, &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(int16(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(int32(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(int64(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(int(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(uint(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(uintptr(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(uint8(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(uint16(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(uint32(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(uint64(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(float32(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(float64(i), &toVal)
	assert.Equal(t, targetI, toVal)

	at.Int64(fmt.Sprintf("%d", i), &toVal)
	assert.Equal(t, targetI, toVal)

	err := at.Int64(uint64(math.MaxInt64)+1, &toVal)
	assert.Equal(t, errors.New("gu.At.Int64() Error: out of range int64"), err)

	var a any
	err = at.Int64(a, &toVal)
	assert.Equal(t, errors.New("gu.At.Int64() Error: unsupported type"), err)
}

func TestAnyUtils(t *testing.T) {
	var toV int
	at.Int(123, &toV)
	assert.Equal(t, int(123), toV)
	err := at.Int(uint64(math.MaxInt)+1, &toV)
	assert.Equal(t, errors.New("gu.At.Int() Error: gu.At.Int64() Error: out of range int64"), err)

	var toUint64V uint64
	at.Uint64(123, &toUint64V)
	assert.Equal(t, uint64(123), toUint64V)

	at.Uint64(uintptr(123123), &toUint64V)
	assert.Equal(t, uint64(123123), toUint64V)

	at.Uint64(float64(1231.1), &toUint64V)
	assert.Equal(t, uint64(1231), toUint64V)
	at.Uint64(float64(1231.9), &toUint64V)
	assert.Equal(t, uint64(1231), toUint64V)
	at.Uint64("1231", &toUint64V)
	assert.Equal(t, uint64(1231), toUint64V)

	err = at.Uint64(-1, &toUint64V)
	assert.Equal(t, errors.New("gu.At.Uint64() Error: out of range uint64"), err)

	var toUintV uint
	at.Uint(123, &toUintV)
	assert.Equal(t, uint(123), toUintV)

	err = at.Uint(-1, &toUintV)
	assert.Equal(t, errors.New("gu.At.Uint() Error: gu.At.Uint64() Error: out of range uint64"), err)

	var toFv float64
	at.Float(321, &toFv)
	assert.Equal(t, float64(321), toFv)

	at.Float(float32(123.3), &toFv)
	assert.Equal(t, float64(float32(123.3)), toFv)

	at.Float(float64(123.3), &toFv)
	assert.Equal(t, float64(123.3), toFv)

	at.Float("123.3", &toFv)
	assert.Equal(t, float64(123.3), toFv)

	var as = []string{"a", "1"}
	err = at.Float(as, &toFv)
	assert.Equal(t, errors.New("gu.At.Float() Error: unsupported type"), err)

	err = at.Float(uint64(math.MaxUint), &toFv)
	assert.Equal(t, errors.New("gu.At.Float() Error: out of range float64"), err)

	uint64Arr := make([]uint64, 0)
	fromAny := []any{1, 2, 3, 4, 5}
	at.Uint64Array(fromAny, &uint64Arr)
	assert.Equal(t, []uint64{1, 2, 3, 4, 5}, uint64Arr)

	int64Arr := make([]int64, 0)
	fromAny2 := []any{-1, 2, -3, 4, -5}
	at.Int64Array(fromAny2, &int64Arr)
	assert.Equal(t, []int64{-1, 2, -3, 4, -5}, int64Arr)

	int64Arr2 := make([]int64, 0)
	fromAny3 := []any{-1, []string{"1"}, -3, 4, -5}
	err = at.Int64Array(fromAny3, &int64Arr2)
	assert.Equal(t, errors.New("gu.At.Int64() Error: unsupported type"), err)
}
