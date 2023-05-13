package types

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ut UintType

func TestUintIf(t *testing.T) {
	v := ut.If(true, 1, 0)
	v2 := ut.If(false, 1, 0)
	assert.Equal(t, uint64(1), v)
	assert.Equal(t, uint64(0), v2)
}

func TestUintSort(t *testing.T) {
	var arr = []uint64{1, 3, 4, 2, 5}
	var arrAsc = []uint64{1, 2, 3, 4, 5}
	var arrDesc = []uint64{5, 4, 3, 2, 1}

	ut.ArrayAsc(arr)
	assert.Equal(t, arrAsc, arr)
	fmt.Println(arr)

	ut.ArrayDesc(arr)
	assert.Equal(t, arrDesc, arr)
	fmt.Println(arr)
}

func TestUintFind(t *testing.T) {
	var arr = []uint64{1, 3, 4, 2, 5, 8, 7, 6, 0, 9}
	var arrAsc = []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var arrDesc = []uint64{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	var item uint64 = 3
	pos := ut.Find(item, arr)
	// 二分查找，返回的是升序后的下标
	assert.Equal(t, 3, pos)

	pos = ut.FindSorted(item, arrAsc, true)
	assert.Equal(t, 3, pos)

	pos = ut.FindSorted(item, arrDesc, false)
	assert.Equal(t, 6, pos)

	pos = ut.LoopFind(item, arr)
	assert.Equal(t, 1, pos)

	pos = ut.BinFind(item, arrAsc, true)
	assert.Equal(t, 3, pos)

	pos = ut.BinFind(item, arrDesc, false)
	assert.Equal(t, 6, pos)

	pos = ut.SortAndBinSearch(item, arr)
	assert.Equal(t, 3, pos)

	fmt.Println(arr)
}

func TestUintConvertTo(t *testing.T) {
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

	var testVal uint64 = 1234
	res := ut.ConvertTo(testVal, v1)
	assert.Equal(t, errors.New("toValue must be a pointer"), res)

	res = ut.ConvertTo(testVal, &v6)
	assert.Equal(t, errors.New("unsupported conversion"), res)

	res = ut.ConvertTo(testVal, &v1)
	assert.Equal(t, fmt.Errorf("out of range for %s", "int8"), res)

	res = ut.ConvertTo(testVal, &v2)
	assert.Nil(t, res)
	assert.Equal(t, int16(testVal), v2)

	testVal = math.MaxInt32 + 1
	res = ut.ConvertTo(testVal, &v3)
	assert.Errorf(t, res, "Out of range for %s", "int32")

	res = ut.ConvertTo(testVal, &v4)
	assert.Nil(t, res)
	assert.Equal(t, int64(testVal), v4)

	res = ut.ConvertTo(testVal, &v5)
	assert.Nil(t, res)
	assert.Equal(t, int(testVal), v5)

	ut.ConvertTo(123, &uv1)
	assert.Equal(t, uint8(123), uv1)

	testVal = math.MaxUint16 + 1
	res = ut.ConvertTo(testVal, &uv1)
	assert.Errorf(t, res, "Out of range for %s", "uint8")

	ut.ConvertTo(123, &uv2)
	assert.Equal(t, uint16(123), uv2)

	res = ut.ConvertTo(testVal, &uv2)
	assert.Errorf(t, res, "Out of range for %s", "uint16")

	res = ut.ConvertTo(testVal, &uv3)
	assert.Nil(t, res)
	assert.Equal(t, uint32(testVal), uv3)

	res = ut.ConvertTo(testVal, &uv4)
	assert.Nil(t, res)
	assert.Equal(t, uint64(testVal), uv4)

	res = ut.ConvertTo(testVal, &uv5)
	assert.Nil(t, res)
	assert.Equal(t, uint(testVal), uv5)

	res = ut.ConvertTo(testVal, &uv6)
	assert.Nil(t, res)
	assert.Equal(t, uintptr(testVal), uv6)
}

func TestUintUtils(t *testing.T) {
	u := uint64(2)
	isOdd := ut.IsOdd(u)
	assert.False(t, isOdd)

	isEven := ut.IsEven(u)
	assert.True(t, isEven)

	inR := ut.InRange(u, 1, 3)
	assert.True(t, inR)

	str := ut.Str(u)
	assert.Equal(t, "2", str)

	min := ut.Min(1, 3, 4, 2, 3, 12312312312, 1231231231)
	assert.Equal(t, uint64(1), min)

	max := ut.Max(1, 3, 4, 2, 3, 213231231, 1231231231)
	assert.Equal(t, uint64(1231231231), max)

	sum := ut.Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	assert.Equal(t, uint64(55), sum)

	avg := ut.Avg(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	assert.Equal(t, 5.5, avg)

	inArr := ut.InArray(u, []uint64{1, 3, 4, 5, 6, 2})
	assert.True(t, inArr)

	inArr2 := ut.InSortedArray(u, []uint64{0, 1, 2, 3, 4, 5}, true)
	assert.True(t, inArr2)
}
