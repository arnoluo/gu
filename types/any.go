package types

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// If 根据条件判断返回不同的值。
func (at AnyType) If(isTrue bool, trueValue, falseValue any) any {
	if isTrue {
		return trueValue
	}

	return falseValue
}

// Int64 将 any 类型的值转换为 int64 类型的值。
// 支持以下类型：uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, string（可以是 10 进制数字字符串）。
// 注：float 类型转换时将损失精度，所以请传入 1111.0 这样不会损失精度的值
func (at AnyType) Int64(fromVal any, toValue *int64) (err error) {
	switch fromVal := fromVal.(type) {
	case int, int8, int16, int32, int64:
		*toValue = reflect.ValueOf(fromVal).Int()
	case uint8, uint16, uint32:
		*toValue = int64(reflect.ValueOf(fromVal).Uint())
	case uint, uintptr:
		v := reflect.ValueOf(fromVal).Uint()
		if v > math.MaxInt64 {
			err = errors.New("gu.At.Int64() Error: out of range int64")
		}
		*toValue = int64(v)
	case uint64:
		if fromVal > math.MaxInt64 {
			err = errors.New("gu.At.Int64() Error: out of range int64")
		}
		*toValue = int64(fromVal)
	case float32, float64:
		*toValue = int64(reflect.ValueOf(fromVal).Float())
	case string:
		*toValue, err = strconv.ParseInt(fromVal, 10, 64)
	default:
		err = errors.New("gu.At.Int64() Error: unsupported type")
	}

	return
}

// Int 将 any 类型的值转换为 int 类型的值。
// 支持以下类型：uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, string（可以是 10 进制数字字符串）。
func (at AnyType) Int(fromVal any, toVal *int) error {
	var to int64
	if err := at.Int64(fromVal, &to); err != nil {
		return fmt.Errorf("gu.At.Int() Error: %s", err.Error())
	}

	*toVal = int(to)
	if to >= math.MinInt && to <= math.MaxInt {
		return nil
	}

	return fmt.Errorf("gu.At.Int() Error: out of range int")
}

// Uint64 将 any 类型的值转换为 uint64 类型的值。
// 支持以下类型：uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, string（可以是 10 进制数字字符串）。
func (at AnyType) Uint64(fromVal any, toValue *uint64) (err error) {
	switch fromVal := fromVal.(type) {
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(fromVal).Int()
		*toValue = uint64(v)
		if v < 0 {
			err = errors.New("gu.At.Uint64() Error: out of range uint64")
		}
	case uint, uintptr, uint8, uint16, uint32, uint64:
		*toValue = reflect.ValueOf(fromVal).Uint()
	case float32, float64:
		v := reflect.ValueOf(fromVal).Float()
		*toValue = uint64(v)
		if v < 0 {
			err = errors.New("gu.At.Uint64() Error: out of range uint64")
		}
	case string:
		*toValue, err = strconv.ParseUint(fromVal, 10, 64)
	default:
		err = errors.New("gu.At.Uint64() Error: unsupported type")
	}

	return
}

// Uint 将 any 类型的值转换为 uint 类型的值。
// 支持以下类型：uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, string（可以是 10 进制数字字符串）。
func (at AnyType) Uint(fromVal any, toVal *uint) error {
	var to uint64
	if err := at.Uint64(fromVal, &to); err != nil {
		return fmt.Errorf("gu.At.Uint() Error: %s", err.Error())
	}

	*toVal = uint(to)
	if to <= math.MaxUint {
		return nil
	}

	return fmt.Errorf("gu.At.Uint() Error: out of range uint")
}

// Float 将 any 类型的值转换为 float64 类型的值。
// 支持以下类型：uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, string（可以是浮点数字符串）。
func (at AnyType) Float(fromVal any, toVal *float64) (err error) {
	switch fromVal := fromVal.(type) {
	case uint8, uint16, uint32:
		*toVal = float64(reflect.ValueOf(fromVal).Uint())
	case uint, uintptr, uint64:
		v := reflect.ValueOf(fromVal).Uint()
		if v > math.Float64bits(math.MaxFloat64) {
			err = errors.New("gu.At.Float() Error: out of range float64")
		}
		*toVal = float64(v)
	case int, int8, int16, int32, int64:
		*toVal = float64(reflect.ValueOf(fromVal).Int())
	case float32:
		*toVal = reflect.ValueOf(fromVal).Float()
	case float64:
		*toVal = fromVal
	case string:
		*toVal, err = strconv.ParseFloat(fromVal, 64)
	default:
		err = errors.New("gu.At.Float() Error: unsupported type")
	}

	return
}

// 结构转换, 使用src内所有kv关系，对dst进行赋值
// src: struct
// dst: struct pointer
func (at AnyType) StructTo(src, dst any) error {
	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst)

	if srcType.Kind() != reflect.Struct {
		return errors.New("src must be struct")
	}
	if dstValue.Kind() != reflect.Ptr {
		return errors.New("dst must be pointer")
	}

	for i := 0; i < srcType.NumField(); i++ {
		dstField := dstValue.Elem().FieldByName(srcType.Field(i).Name)
		if dstField.CanSet() {
			dstField.Set(srcValue.Field(i))
		}
	}

	return nil
}

// Return true if stack has the element item, return false otherwise
// 由于使用了反射，可能在性能上有一定的负担，并且由于没有类型检查，不能有效地避免类型不匹配的情况
// 不推荐使用这种方式，请转换为具体类型后执行类型下的 InArray
func (at AnyType) InArray(item, stack any) bool {
	arrType := reflect.TypeOf(stack)
	kd := arrType.Kind()
	if kd == reflect.Slice || kd == reflect.Array {
		v := reflect.ValueOf(stack)
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Interface() == item {
				return true
			}
		}
	}

	return false
}

// any 数组转为 uint64 数组
func (at AnyType) Uint64Array(arr []any, dstArr *[]uint64) error {
	for _, v := range arr {
		var toVal uint64
		if err := at.Uint64(v, &toVal); err != nil {
			return err
		}
		*dstArr = append(*dstArr, toVal)
	}

	return nil
}

// any 数组转为 int64 数组
func (at AnyType) Int64Array(arr []any, dstArr *[]int64) error {
	for _, v := range arr {
		var toVal int64
		if err := at.Int64(v, &toVal); err != nil {
			return err
		}
		*dstArr = append(*dstArr, toVal)
	}

	return nil
}
