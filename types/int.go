package types

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
)

// If 根据条件判断返回不同的值。
func (it IntType) If(isTrue bool, trueValue, falseValue int64) int64 {
	if isTrue {
		return trueValue
	}

	return falseValue
}

// 将 int64 类型转换为 其他int 类型， 仅限int类型内转换
//
// toVal 必须为 int相关类型的指针 eg.
//
//	var val int8
//	if err := ConvertTo(12, &val); err != nil {}
func (it IntType) ConvertTo(from int64, toVal any) error {
	rv := reflect.ValueOf(toVal)
	if rv.Kind() != reflect.Ptr {
		return errors.New("toVal must be a pointer")
	}

	val := rv.Elem() // dereference pointer to get actual value
	if !val.IsValid() {
		return errors.New("toVal must not be nil")
	}

	var hasErr, isUint bool
	switch val.Kind() {
	case reflect.Int:
		if !it.InIntRange(from) {
			hasErr = true
		}
	case reflect.Int8:
		if !it.InInt8Range(from) {
			hasErr = true
		}
	case reflect.Int16:
		if !it.InInt16Range(from) {
			hasErr = true
		}
	case reflect.Int32:
		if !it.InInt32Range(from) {
			hasErr = true
		}
	case reflect.Int64:
	case reflect.Uint, reflect.Uintptr:
		isUint = true
		if !it.InUintRange(from) {
			hasErr = true
		}
	case reflect.Uint8:
		isUint = true
		if !it.InUint8Range(from) {
			hasErr = true
		}
	case reflect.Uint16:
		isUint = true
		if !it.InUint16Range(from) {
			hasErr = true
		}
	case reflect.Uint32:
		isUint = true
		if !it.InUint32Range(from) {
			hasErr = true
		}
	case reflect.Uint64:
		isUint = true
		if !it.InUint64Range(from) {
			hasErr = true
		}
	default:
		return errors.New("unsupported conversion")
	}

	if hasErr {
		return fmt.Errorf("out of range for %s", val.Type().Name())
	}

	if isUint {
		val.SetUint(uint64(from))
	} else {
		val.SetInt(from)
	}

	return nil
}

// 判断给定的整数是否在 int8 范围内
func (it IntType) InInt8Range(val int64) bool {
	return it.InRange(val, math.MinInt8, math.MaxInt8)
}

// 判断给定的整数是否在 int16 范围内
func (it IntType) InInt16Range(val int64) bool {
	return it.InRange(val, math.MinInt16, math.MaxInt16)
}

// 判断给定的整数是否在 int32 范围内
func (it IntType) InInt32Range(val int64) bool {
	return it.InRange(val, math.MinInt32, math.MaxInt32)
}

// 判断给定的整数是否在 int 范围内
func (it IntType) InIntRange(val int64) bool {
	return it.InRange(val, math.MinInt, math.MaxInt)
}

// 判断给定的整数是否在 uint8 范围内
func (it IntType) InUint8Range(val int64) bool {
	return it.InRange(val, 0, math.MaxUint8)
}

// 判断给定的整数是否在 uint16 范围内
func (it IntType) InUint16Range(val int64) bool {
	return it.InRange(val, 0, math.MaxUint16)
}

// 判断给定的整数是否在 uint32 范围内
func (it IntType) InUint32Range(val int64) bool {
	return it.InRange(val, 0, math.MaxUint32)
}

// 判断给定的整数是否在 uint 范围内
func (it IntType) InUintRange(val int64) bool {
	return val >= 0 && uint64(val) <= math.MaxUint
}

// 判断给定的整数是否在 uint64 范围内
func (it IntType) InUint64Range(val int64) bool {
	return val >= 0
}

// 实现判断指定的 int64 类型值是否为奇数的功能，返回布尔值
func (it IntType) IsOdd(value int64) bool {
	return !it.IsEven(value)
}

// 实现判断指定的 int64 类型值是否为偶数的功能，返回布尔值
func (it IntType) IsEven(value int64) bool {
	return value%2 == 0
}

// 实现判断指定的 int64 类型值是否在指定的范围内的功能，返回布尔值
func (it IntType) InRange(value, min, max int64) bool {
	return value >= min && value <= max
}

// Str：将 int64 类型的值转换为字符串类型的值。
func (it IntType) Str(value int64) string {
	return strconv.FormatInt(value, 10)
}

// Abs：获取 int64 类型的值的绝对值。
func (it IntType) Abs(value int64) int64 {
	if value < 0 {
		return -value
	}
	return value
}

// Min：获取 int64 类型的值数组中的最小值。
func (it IntType) Min(values ...int64) int64 {
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

// Max：获取 int64 类型的值数组中的最大值。
func (it IntType) Max(values ...int64) int64 {
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

// Sum：计算 int64 类型的值数组中所有值的总和。
func (it IntType) Sum(values ...int64) int64 {
	var sum int64 = 0
	for _, v := range values {
		sum += v
	}
	return sum
}

// Avg：计算 int64 类型的值数组中所有值的平均值。
func (it IntType) Avg(values ...int64) float64 {
	if len(values) == 0 {
		return 0 // 或返回一个错误
	}
	sum := it.Sum(values...)
	return float64(sum) / float64(len(values))
}

// 整数切片查找，threshold 参数生效
func (it IntType) InArray(value int64, arr []int64) bool {
	return it.Find(value, arr) >= 0
}

// 已排序数组查找，threshold 参数生效
func (it IntType) InSortedArray(value int64, arr []int64, isAsc bool) bool {
	return it.FindSorted(value, arr, isAsc) >= 0
}

// 遍历查找数组
//
// 如果数组长度较长或对同一数组做多次 LoopFind，建议先 ArrayAsc 后使用 BinFind
//
// 成功时返回查找到的数组下标，失败返回 -1
func (it IntType) LoopFind(value int64, arr []int64) int {
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
func (it IntType) BinFind(value int64, arr []int64, isAsc bool) int {
	return it.BinSearch(value, arr, isAsc)
}

// 查找数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找
//
// 二分查找时，将会对数组进行升序排序，查找成功返回的也会是升序后的下标
//
// 成功时返回查找到的数组下标，失败返回 -1
func (it IntType) Find(value int64, arr []int64) int {
	if len(arr) > threshold {
		return it.SortAndBinSearch(value, arr)
	} else {
		// 否则使用遍历查找
		return it.LoopFind(value, arr)
	}
}

// 查找已排序数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找
//
// 成功时返回查找到的数组下标，失败返回 -1
func (it IntType) FindSorted(value int64, arr []int64, isAsc bool) int {
	if len(arr) > threshold {
		return it.BinFind(value, arr, isAsc)
	} else {
		// 否则使用遍历查找
		return it.LoopFind(value, arr)
	}
}

// 二分查找(此函数不会进行排序，请输入已排序的数组，否则会产生非预期结果)
// 成功返回查找到的数组下标，失败返回 -1
func (it IntType) BinSearch(value int64, arr []int64, isAsc bool) int {
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
func (it IntType) ArrayAsc(arr []int64) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

// 数组排序 desc
func (it IntType) ArrayDesc(arr []int64) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
}

// 对数组排序并进行二分查找法，成功返回查找到的数组下标，失败返回 -1
func (it IntType) SortAndBinSearch(value int64, arr []int64) int {
	tmp := make([]int64, len(arr))
	copy(tmp, arr)
	it.ArrayAsc(tmp)
	return it.BinSearch(value, tmp, true)
}
