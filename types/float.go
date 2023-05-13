package types

import (
	"math"
	"sort"
	"strconv"
)

// If 根据条件判断返回不同的值。
func (ft FloatType) If(isTrue bool, trueValue, falseValue float64) float64 {
	if isTrue {
		return trueValue
	}
	return falseValue
}

// Str 将浮点类型的值转换为字符串类型的值。
func (ft FloatType) Str(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

// Abs 返回浮点数的绝对值。
func (ft FloatType) Abs(value float64) float64 {
	return math.Abs(value)
}

// Ceil 返回不小于 value 的最小整数，也就是向上取整。
func (ft FloatType) Ceil(value float64) float64 {
	return math.Ceil(value)
}

// Floor 返回不大于 value 的最大整数，也就是向下取整。
func (ft FloatType) Floor(value float64) float64 {
	return math.Floor(value)
}

// Round 对浮点数进行四舍五入操作，places 参数表示小数点后保留位数。
func (ft FloatType) Round(value float64, places int) float64 {
	precision := math.Pow(10, float64(places))
	return math.Round(value*precision) / precision
}

// Min 获取浮点数数组中的最小值。
func (ft FloatType) Min(values ...float64) float64 {
	if len(values) == 0 {
		return 0 // 或返回一个错误
	}
	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// Max 获取浮点数数组中的最大值。
func (ft FloatType) Max(values ...float64) float64 {
	if len(values) == 0 {
		return 0 // 或返回一个错误
	}
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// Sum 计算浮点数数组中所有值的总和。
func (ft FloatType) Sum(values ...float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum
}

// Avg 计算浮点数数组中所有值的平均值。
func (ft FloatType) Avg(values ...float64) float64 {
	if len(values) == 0 {
		return 0 // 或返回一个错误
	}
	sum := ft.Sum(values...)
	return sum / float64(len(values))
}

// 浮点数切片查找，threshold 参数生效
func (ft FloatType) InArray(value float64, arr []float64) bool {
	return ft.Find(value, arr) >= 0
}

// 已排序数组查找，threshold 参数生效
func (ft FloatType) InSortedArray(value float64, arr []float64, isAsc bool) bool {
	return ft.FindSorted(value, arr, isAsc) >= 0
}

// 遍历查找数组
//
// 如果数组长度较长或对同一数组做多次 LoopFind，建议先 ArrayAsc 后使用 BinFind
//
// 成功时返回查找到的数组下标，失败返回 -1
func (ft FloatType) LoopFind(value float64, arr []float64) int {
	for i, a := range arr {
		if a == value {
			return i
		}
	}
	return -1
}

// 二分查找数组
//
// 此函数适用于数组值已排序或将对同一数组进行多次查找，传入已排序的数组以提高效率。
// 注：若传入未排序的数组，结果可能并不符合预期。
//
// 成功时返回查找到的数组下标，失败返回 -1
func (ft FloatType) BinFind(value float64, arr []float64, isAsc bool) int {
	return ft.BinSearch(value, arr, isAsc)
}

// 查找数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找
//
// 二分查找时，将会对数组进行升序排序，查找成功返回的也会是升序后的下标
//
// 成功时返回查找到的数组下标，失败返回 -1
func (ft FloatType) Find(value float64, arr []float64) int {
	if len(arr) > threshold {
		return ft.SortAndBinSearch(value, arr)
	} else {
		// 否则使用遍历查找
		return ft.LoopFind(value, arr)
	}
}

// 查找已排序数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找
//
// 成功时返回查找到的数组下标，失败返回 -1
func (ft FloatType) FindSorted(value float64, arr []float64, isAsc bool) int {
	if len(arr) > threshold {
		return ft.BinFind(value, arr, isAsc)
	} else {
		// 否则使用遍历查找
		return ft.LoopFind(value, arr)
	}
}

// 二分查找(此函数不会进行排序，请输入已排序的数组，否则会产生非预期结果)
// 成功返回查找到的数组下标，失败返回 -1
func (ft FloatType) BinSearch(value float64, arr []float64, isAsc bool) int {
	l, r := 0, len(arr)-1
	if isAsc {
		for l <= r {
			mid := (l + r) / 2
			if arr[mid] < value {
				l = mid + 1
			} else if arr[mid] > value {
				r = mid - 1
			} else {
				return mid
			}
		}
	} else {
		for l <= r {
			mid := (l + r) / 2
			if arr[mid] > value {
				l = mid + 1
			} else if arr[mid] < value {
				r = mid - 1
			} else {
				return mid
			}
		}
	}

	return -1
}

// 数组排序 asc
func (ft FloatType) ArrayAsc(arr []float64) {
	sort.Float64s(arr)
}

// 数组排序 desc
func (ft FloatType) ArrayDesc(arr []float64) {
	sort.Sort(sort.Reverse(sort.Float64Slice(arr)))
}

// 对数组排序并进行二分查找法，成功返回查找到的数组下标，失败返回 -1
func (ft FloatType) SortAndBinSearch(value float64, arr []float64) int {
	tmp := make([]float64, len(arr))
	copy(tmp, arr)
	ft.ArrayAsc(tmp)
	return ft.BinSearch(value, tmp, true)
}
