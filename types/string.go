package types

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	// 字母 + 数字，移除 il o0
	CHARS string = `abcdefghjkmnpqrstuvwxyz123456789`

	// 纯字母，移除 il
	PURE_LETTER_CHARS string = `abcdefghjkmnopqrstuvwxyz`

	// 纯数字字符
	PURE_NUMBER_CHARS string = `0123456789`
)

// string 转 int, 需设置转换错误时的默认值
func (st StrType) Int(str string, errValue int) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		return errValue
	}
	return value
}

// string 转 bool, 需设置转换错误时的默认值
func (st StrType) Bool(str string, errBool bool) bool {
	value, err := strconv.ParseBool(str)
	if err != nil {
		return errBool
	}
	return value
}

// Covert string to unsigned int using strconv.Atoi(), return errValue if err != nil or value < 0
func (st StrType) Uint(str string, errValue uint) uint {
	v, err := strconv.ParseUint(str, 10, 0)

	if err != nil {
		return errValue
	}

	return uint(v)
}

// Covert string to positive int using strconv.Atoi(), return errValue if err != nil or value <= 0
func (st StrType) Pint(str string, errValue uint) uint {
	value := st.Uint(str, errValue)
	if value <= 0 {
		return errValue
	}

	return value
}

// Covert string to positive int using strconv.Atoi(), return errValue if err != nil or value <= 0
func (st StrType) Float(str string, errValue float64) float64 {
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return errValue
	}

	return v
}

func (st StrType) IsInt(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func (st StrType) IsNum(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func (st StrType) IsEmpty(value string) bool {
	return value == ""
}

// utf8(6 bytes at most) substring
func (st StrType) Sub(str string, begin, length int) string {
	rs := []rune(str)
	lth := len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}

	end := begin + length

	if end > lth {
		end = lth
	}

	return string(rs[begin:end])
}

// Return string param trueValue if boolValue=true, return string param falseValue otherwise
func (st StrType) If(boolValue bool, trueValue, falseValue string) string {
	if boolValue {
		return trueValue
	}

	return falseValue
}

// Replace baseStr with regexpPattern to replacement
func (st StrType) RegReplace(baseStr, regexpPattern, replacement string) string {
	reg := regexp.MustCompile(regexpPattern)
	return reg.ReplaceAllString(baseStr, replacement)
}

// Generate random str base on letter&number mixed chars
func (st StrType) Rand(length int) string {
	return st.RandChars(CHARS, length)
}

// Generate random str base on letter chars
func (st StrType) RandLetters(length int) string {
	return st.RandChars(PURE_LETTER_CHARS, length)
}

// Generate random str base on number chars
func (st StrType) RandNumbers(length int) string {
	return st.RandChars(PURE_NUMBER_CHARS, length)
}

// Generate random str base on baseChars
func (st StrType) RandChars(chars string, length int) (str string) {
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		randPos := rand.Intn(len(chars))
		str += chars[randPos : randPos+1]
	}
	return
}

// 返回字符串的长度
func (st StrType) Len(str string) int {
	return len(str)
}

// 将字符串首字母大写
func (st StrType) UpperFirst(str string) string {
	if str == "" {
		return ""
	}

	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// 将字符串首字母小写
func (st StrType) LowerFirst(str string) string {
	if str == "" {
		return ""
	}

	runes := []rune(str)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// 将字符串首尾的空白字符去除
func (st StrType) TrimSpace(str string) string {
	return strings.TrimSpace(str)
}

// 将字符串左右两侧指定字符集合 charsets 中的字符去除
func (st StrType) Trim(str, charsets string) string {
	return strings.Trim(str, charsets)
}

// 将字符串左侧指定字符集合 charsets 中的字符去除
func (st StrType) Ltrim(str, charsets string) string {
	return strings.TrimLeft(str, charsets)
}

// 将字符串右侧指定字符集合 charsets 中的字符去除
func (st StrType) Rtrim(str, charsets string) string {
	return strings.TrimRight(str, charsets)
}

// 以 sep 为分隔符拼接字符串数组为一个字符串，同 strings.Join()
func (st StrType) Join(strs []string, sep string) string {
	return strings.Join(strs, sep)
}

// 将字符串 str 照sep进行分割，并返回分割后的字符串数组，同 strings.Split()
func (st StrType) Split(str, sep string) []string {
	return strings.Split(str, sep)
}

// 查找字符串 substr 在 str 中首次出现的位置，如果找不到返回 -1
func (st StrType) Index(str, substr string) int {
	return strings.Index(str, substr)
}

// 替换字符串中的 old 为 new，n 为替换的最大次数（小于等于 0 表示全部替换）
func (st StrType) Replace(str, old, new string, n int) string {
	return strings.Replace(str, old, new, n)
}

// 判断字符串 str 是否以 prefix 开头
func (st StrType) HasPrefix(str, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

// 判断字符串 str 是否以 suffix 结尾
func (st StrType) HasSuffix(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}

// 格式化字符串，类似 fmt.Sprintf()
func (st StrType) Format(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

// 字符串切片查找，threshold 参数生效
func (st StrType) InArray(v string, arr []string) bool {
	return st.Find(v, arr) >= 0
}

// 查找数组，threshold 参数生效，按照数组总量决定查找策略(>= 8 二分查找；否则遍历查找)
// 成功时返回查找到的数组下标，失败返回 -1
func (st StrType) Find(v string, arr []string) int {
	if len(arr) < threshold {
		for i, a := range arr {
			if a == v {
				return i
			}
		}
		return -1
	} else {
		return st.BinSearch(v, arr)
	}
}

// 二分查找法，成功返回查找到的数组下标，失败返回 -1
func (st StrType) BinSearch(v string, arr []string) int {
	sort.Strings(arr)
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
