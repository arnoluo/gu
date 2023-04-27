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
func (ut UintType) If(isTrue bool, trueValue, falseValue uint64) uint64 {
	if isTrue {
		return trueValue
	}

	return falseValue
}

// 将 uint64 类型转换为 其他int 类型， 仅限int类型内转换
//
// toVal 必须为 int相关类型的指针 eg.
//
//	var val int8
//	if err := ConvertTo(12, &val); err != nil {}
func (ut UintType) ConvertTo(from uint64, toVal any) error {
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
		if from > math.MaxInt {
			hasErr = true
		}
	case reflect.Int8:
		if from > math.MaxInt8 {
			hasErr = true
		}
	case reflect.Int16:
		if from > math.MaxInt16 {
			hasErr = true
		}
	case reflect.Int32:
		if from > math.MaxInt32 {
			hasErr = true
		}
	case reflect.Int64:
		if from > math.MaxInt64 {
			hasErr = true
		}
	case reflect.Uint, reflect.Uintptr:
		isUint = true
		if from > math.MaxUint {
			hasErr = true
		}
	case reflect.Uint8:
		isUint = true
		if from > math.MaxUint8 {
			hasErr = true
		}
	case reflect.Uint16:
		isUint = true
		if from > math.MaxUint16 {
			hasErr = true
		}
	case reflect.Uint32:
		isUint = true
		if from > math.MaxUint32 {
			hasErr = true
		}
	case reflect.Uint64:
		isUint = true
	default:
		return errors.New("unsupported conversion")
	}

	if hasErr {
		return fmt.Errorf("out of range for %s", val.Type().Name())
	}

	if isUint {
		val.SetUint(from)
	} else {
		val.SetInt(int64(from))
	}

	return nil
}

// 实现判断指定的 uint64 类型值是否为奇数的功能，返回布尔值
func (ut UintType) IsOdd(value uint64) bool {
	return !ut.IsEven(value)
}

// 实现判断指定的 uint64 类型值是否为偶数的功能，返回布尔值
func (ut UintType) IsEven(value uint64) bool {
	return value%2 == 0
}

// 实现判断指定的 uint64 类型值是否在指定的范围内的功能，返回布尔值
func (ut UintType) InRange(value, min, max uint64) bool {
	return value >= min && value <= max
}

// 将 uint64 类型的数据转化为字符串类型
func (ut UintType) Str(value uint64) string {
	return strconv.FormatUint(value, 10)
}

// Min：获取 uint64 类型的值数组中的最小值。
func (ut UintType) Min(values ...uint64) uint64 {
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

// Max：获取 uint64 类型的值数组中的最大值。
func (ut UintType) Max(values ...uint64) uint64 {
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

// Sum：计算 uint64 类型的值数组中所有值的总和。
func (ut UintType) Sum(values ...uint64) uint64 {
	var sum uint64 = 0
	for _, v := range values {
		sum += v
	}
	return sum
}

// Avg：计算 uint64 类型的值数组中所有值的平均值。
func (ut UintType) Avg(values ...uint64) float64 {
	if len(values) == 0 {
		return 0 // 或返回一个错误
	}
	sum := ut.Sum(values...)
	return float64(sum) / float64(len(values))
}

// 整数切片查找，threshold 参数生效
func (ut UintType) InArray(v uint64, arr []uint64) bool {
	return ut.Find(v, arr) >= 0
}

// 查找数组，threshold 参数生效，按照数组总量决定查找策略(>= 8 二分查找；否则遍历查找)
//
// 二分查找将会对数组做值排序，所以如果对同一数组做多次 Find，建议先 ArrayAsc 后使用 FindSorted
//
// 成功时返回查找到的数组下标，失败返回 -1
func (ut UintType) Find(v uint64, arr []uint64) int {
	if len(arr) < threshold {
		for i, a := range arr {
			if a == v {
				return i
			}
		}
		return -1
	} else {
		return ut.SortAndBinSearch(v, arr)
	}
}

// 查找已排序数组，threshold 参数生效，按照数组总量决定查找策略(>= 8 二分查找；否则遍历查找)
//
// 此函数适用于数组值已排序或将对同一数组进行多次查找，传入已排序的数组以提高效率。
// 注：若传入未排序的数组，结果可能并不符合预期。
//
// 成功时返回查找到的数组下标，失败返回 -1
func (ut UintType) FindSorted(v uint64, arr []uint64) int {
	if len(arr) < threshold {
		for i, a := range arr {
			if a == v {
				return i
			}
		}
		return -1
	} else {
		return ut.BinSearch(v, arr)
	}
}

// 数组排序 asc
func (ut UintType) ArrayAsc(arr []uint64) []uint64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

// 数组排序 desc
func (ut UintType) ArrayDesc(arr []uint64) []uint64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	return arr
}

// 二分查找法(此函数不会进行排序，请输入已排序的数组)，成功返回查找到的数组下标，失败返回 -1
func (ut UintType) BinSearch(v uint64, arr []uint64) int {

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

// 对数组排序并进行二分查找法，成功返回查找到的数组下标，失败返回 -1
func (ut UintType) SortAndBinSearch(v uint64, arr []uint64) int {
	return ut.BinSearch(v, ut.ArrayAsc(arr))
}
