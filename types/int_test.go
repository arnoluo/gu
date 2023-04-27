package types

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var it IntType

func TestIntIf(t *testing.T) {
	assert.Equal(t, int64(1), it.If(true, 1, 0))
	assert.Equal(t, int64(0), it.If(false, 1, 0))
}

func TestConvertTo(t *testing.T) {
	var v1 int8
	var v2 int16
	var v3 int32
	var v4 int64
	var v5 int
	var v6 any

	var uv1 uint8
	var uv2 uint16
	var uv3 uint32
	var uv4 uint64
	var uv5 uint
	var uv6 uintptr

	var testVal int64 = 1234
	res := it.ConvertTo(testVal, v1)
	assert.Errorf(t, res, "toVal must be a pointer")

	res = it.ConvertTo(testVal, &v6)
	assert.Errorf(t, res, "toVal must not be nil")

	res = it.ConvertTo(testVal, &v1)
	assert.Errorf(t, res, "Out of range for %s", "int8")

	res = it.ConvertTo(testVal, &v2)
	assert.Nil(t, res)
	assert.Equal(t, int16(testVal), v2)

	testVal = math.MaxInt32 + 1
	res = it.ConvertTo(testVal, &v3)
	assert.Errorf(t, res, "Out of range for %s", "int32")

	res = it.ConvertTo(testVal, &v4)
	assert.Nil(t, res)
	assert.Equal(t, int64(testVal), v4)

	res = it.ConvertTo(testVal, &v5)
	assert.Nil(t, res)
	assert.Equal(t, int(testVal), v5)

	testVal = math.MaxUint16 + 1
	res = it.ConvertTo(testVal, uv1)
	assert.Errorf(t, res, "Out of range for %s", "uint8")

	res = it.ConvertTo(testVal, uv2)
	assert.Errorf(t, res, "Out of range for %s", "uint16")

	res = it.ConvertTo(testVal, &uv3)
	assert.Nil(t, res)
	assert.Equal(t, uint32(testVal), uv3)

	testVal = -1
	res = it.ConvertTo(testVal, &uv4)
	assert.Errorf(t, res, "Out of range for %s", "uint64")

	res = it.ConvertTo(testVal, &uv5)
	assert.Errorf(t, res, "Out of range for %s", "uint")

	res = it.ConvertTo(testVal, &uv6)
	assert.Errorf(t, res, "Out of range for %s", "uintptr")
}

func TestIntStr(t *testing.T) {
	var i IntType
	assert.Equal(t, "1", i.Str(1))
}
