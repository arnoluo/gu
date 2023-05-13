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
// toValue 必须为 int相关类型的指针 eg.
//
//	var val int8
//	if err := ConvertTo(12, &val); err != nil {}
func (ut UintType) ConvertTo(from uint64, toValue any) error {
	rv := reflect.ValueOf(toValue)
	if rv.Kind() != reflect.Ptr {
		return errors.New("toValue must be a pointer")
	}

	val := rv.Elem() // dereference pointer to get actual value

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
	for _, value := range values[1:] {
		if value < min {
			min = value
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
	for _, value := range values[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// Sum：计算 uint64 类型的值数组中所有值的总和。
func (ut UintType) Sum(values ...uint64) uint64 {
	var sum uint64 = 0
	for _, value := range values {
		sum += value
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
func (ut UintType) InArray(value uint64, arr []uint64) bool {
	return ut.Find(value, arr) >= 0
}

// 已排序数组查找，threshold 参数生效
func (ut UintType) InSortedArray(value uint64, arr []uint64, isAsc bool) bool {
	return ut.FindSorted(value, arr, isAsc) >= 0
}

// 遍历查找数组
//
// 如果数组长度较长或对同一数组做多次 LoopFind，建议先 ArrayAsc 后使用 BinFind
//
// 成功时返回查找到的数组下标，失败返回 -1
func (ut UintType) LoopFind(value uint64, arr []uint64) int {
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
func (ut UintType) BinFind(value uint64, arr []uint64, isAsc bool) int {
	return ut.BinSearch(value, arr, isAsc)
}

// 查找数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找
//
// 二分查找时，将会对数组进行升序排序，查找成功返回的也会是升序后的下标
//
// 成功时返回查找到的数组下标，失败返回 -1
func (ut UintType) Find(value uint64, arr []uint64) int {
	if len(arr) > threshold {
		return ut.SortAndBinSearch(value, arr)
	} else {
		// 否则使用遍历查找
		return ut.LoopFind(value, arr)
	}
}

// 查找已排序数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找
//
// 成功时返回查找到的数组下标，失败返回 -1
func (ut UintType) FindSorted(value uint64, arr []uint64, isAsc bool) int {
	if len(arr) > threshold {
		return ut.BinFind(value, arr, isAsc)
	} else {
		// 否则使用遍历查找
		return ut.LoopFind(value, arr)
	}
}

// 数组排序 asc
func (ut UintType) ArrayAsc(arr []uint64) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

// 数组排序 desc
func (ut UintType) ArrayDesc(arr []uint64) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
}

// 二分查找(此函数不会进行排序，请输入已排序的数组，否则会产生非预期结果)
// 成功返回查找到的数组下标，失败返回 -1
func (ut UintType) BinSearch(value uint64, arr []uint64, isAsc bool) int {
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

// 对数组排序(不会改变原数组顺序)并进行二分查找法，成功返回查找到的数组下标，失败返回 -1
func (ut UintType) SortAndBinSearch(value uint64, arr []uint64) int {
	tmp := make([]uint64, len(arr))
	copy(tmp, arr)
	ut.ArrayAsc(tmp)
	return ut.BinSearch(value, tmp, true)
}
