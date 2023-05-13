package gu

import (
	"os"

	"github.com/arnoluo/gu/types"
)

var (
	// StrType
	St types.StrType

	// IntType
	It types.IntType

	// UintType
	Ut types.UintType

	// BoolType
	Bt types.BoolType

	// FloatType
	Ft types.FloatType

	// AnyType
	At types.AnyType

	// // ListType
	// Lt types.ListType

	// // RingType
	// Rt types.RingType

	// // HeapType
	// Ht types.HeapType
)

// 获取环境变量
func Env(name string) string {
	return os.Getenv(name)
}

// 获取环境变量，返回int, 发生错误时返回指定默认值
func EnvInt(name string, defaultValue int) int {
	return St.Int(Env(name), defaultValue)
}
