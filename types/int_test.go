package types

import (
	"fmt"
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

	it.ConvertTo(123, &uv1)
	assert.Equal(t, uint8(123), uv1)

	testVal = math.MaxUint16 + 1
	res = it.ConvertTo(testVal, &uv1)
	assert.Errorf(t, res, "Out of range for %s", "uint8")

	it.ConvertTo(123, &uv2)
	assert.Equal(t, uint16(123), uv2)

	res = it.ConvertTo(testVal, &uv2)
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

func TestIntFind(t *testing.T) {
	var arr = []int64{1, -3, -2, 2, 3, 4, -1, -5, 0, 5, -4}
	var arrAsc = []int64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	var arrDesc = []int64{5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5}

	var item int64 = -2
	pos := it.Find(item, arr)
	// 二分查找，返回的是升序后的下标
	assert.Equal(t, 3, pos)

	pos = it.FindSorted(item, arrAsc, true)
	assert.Equal(t, 3, pos)

	pos = it.FindSorted(item, arrDesc, false)
	assert.Equal(t, 7, pos)

	pos = it.LoopFind(item, arr)
	assert.Equal(t, 2, pos)

	pos = it.BinFind(item, arrAsc, true)
	assert.Equal(t, 3, pos)

	pos = it.BinFind(item, arrDesc, false)
	assert.Equal(t, 7, pos)

	pos = it.SortAndBinSearch(item, arr)
	assert.Equal(t, 3, pos)

	fmt.Println(arr)
}

func TestIntUtils(t *testing.T) {
	i := int64(2)
	isOdd := it.IsOdd(i)
	assert.False(t, isOdd)

	isEven := it.IsEven(i)
	assert.True(t, isEven)

	inR := it.InRange(i, 1, 3)
	assert.True(t, inR)

	str := it.Str(i)
	assert.Equal(t, "2", str)

	min := it.Min(1, 3, 4, 2, 3, 12312312312, 1231231231)
	assert.Equal(t, int64(1), min)

	max := it.Max(1, 3, 4, 2, 3, 213231231, 1231231231)
	assert.Equal(t, int64(1231231231), max)

	sum := it.Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	assert.Equal(t, int64(55), sum)

	avg := it.Avg(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	assert.Equal(t, 5.5, avg)

	inArr := it.InArray(i, []int64{1, 3, 4, 5, 6, 2})
	assert.True(t, inArr)

	inArr2 := it.InSortedArray(i, []int64{0, 1, 2, 3, 4, 5}, true)
	assert.True(t, inArr2)

	assert.Equal(t, int64(1), it.Abs(-1))
	assert.Equal(t, int64(11), it.Abs(11))
}

func TestIntSort(t *testing.T) {
	var arr = []int64{-1, 3, 4, 2, 5}
	var arrAsc = []int64{-1, 2, 3, 4, 5}
	var arrDesc = []int64{5, 4, 3, 2, -1}

	it.ArrayAsc(arr)
	assert.Equal(t, arrAsc, arr)
	fmt.Println(arr)

	it.ArrayDesc(arr)
	assert.Equal(t, arrDesc, arr)
	fmt.Println(arr)
}
