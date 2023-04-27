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
func (ft FloatType) InArray(v float64, arr []float64) bool {
	return ft.Find(v, arr) >= 0
}

// 查找数组，threshold 参数生效，按照数组总量决定查找策略(>= 8 二分查找；否则遍历查找)
// 成功时返回查找到的数组下标，失败返回 -1
func (ft FloatType) Find(v float64, arr []float64) int {
	if len(arr) < threshold {
		for i, a := range arr {
			if a == v {
				return i
			}
		}
		return -1
	} else {
		return ft.BinSearch(v, arr)
	}
}

// 二分查找法，成功返回查找到的数组下标，失败返回 -1
func (ft FloatType) BinSearch(v float64, arr []float64) int {
	sort.Float64s(arr)
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] < v {
			l = mid + 1
		} else if arr[mid] > v {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
