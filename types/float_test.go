package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ft FloatType

func TestFloatIf(t *testing.T) {
	assert.Equal(t, 1.11, ft.If(true, 1.11, 0.00))
	assert.Equal(t, 0.00, ft.If(false, 1.11, 0.00))
}

func TestFloatFind(t *testing.T) {
	var arr = []float64{1.2, -3.2, -2.1, 2.3, 3.4, 4.5, -1.0, -5.4, 0.00, 5.6, -4.3}
	var arrAsc = []float64{-5.4, -4.3, -3.2, -2.1, -1.0, 0.00, 1.2, 2.3, 3.4, 4.5, 5.6}
	var arrDesc = []float64{5.6, 4.5, 3.4, 2.3, 1.2, 0.00, -1.0, -2.1, -3.2, -4.3, -5.4}

	var item float64 = -2.1
	pos := ft.Find(item, arr)
	// 二分查找，返回的是升序后的下标
	assert.Equal(t, 3, pos)

	pos = ft.FindSorted(item, arrAsc, true)
	assert.Equal(t, 3, pos)

	pos = ft.FindSorted(item, arrDesc, false)
	assert.Equal(t, 7, pos)

	pos = ft.LoopFind(item, arr)
	assert.Equal(t, 2, pos)

	pos = ft.BinFind(item, arrAsc, true)
	assert.Equal(t, 3, pos)

	pos = ft.BinFind(item, arrDesc, false)
	assert.Equal(t, 7, pos)

	pos = ft.SortAndBinSearch(item, arr)
	assert.Equal(t, 3, pos)

	fmt.Println(arr)
}

func TestFloatUtils(t *testing.T) {
	f := float64(1.234)
	assert.Equal(t, "1.234", ft.Str(f))
	assert.Equal(t, float64(1.234), ft.Abs(f*-1))
	assert.Equal(t, float64(2), ft.Ceil(f))
	assert.Equal(t, float64(1), ft.Floor(f))
	assert.Equal(t, float64(1.23), ft.Round(f, 2))

	min := ft.Min(1, 3, 4, 2, 3, 12312312312, 1231231231)
	assert.Equal(t, float64(1), min)

	max := ft.Max(1, 3, 4, 2, 3, 213231231, 1231231231)
	assert.Equal(t, float64(1231231231), max)

	sum := ft.Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	assert.Equal(t, float64(55), sum)

	avg := ft.Avg(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	assert.Equal(t, float64(5.5), avg)

	inArr := ft.InArray(f, []float64{1.234, 3.1, 4, 5, 6, 2})
	assert.True(t, inArr)

	inArr2 := ft.InSortedArray(f, []float64{0, 1.234, 2, 3, 4, 5}, true)
	assert.True(t, inArr2)
}

func TestFloatSort(t *testing.T) {
	var arr = []float64{-1, 3, 4, 2, 5}
	var arrAsc = []float64{-1, 2, 3, 4, 5}
	var arrDesc = []float64{5, 4, 3, 2, -1}

	ft.ArrayAsc(arr)
	assert.Equal(t, arrAsc, arr)
	fmt.Println(arr)

	ft.ArrayDesc(arr)
	assert.Equal(t, arrDesc, arr)
	fmt.Println(arr)
}
