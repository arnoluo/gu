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
	value, err := strconv.ParseUint(str, 10, 0)

	if err != nil {
		return errValue
	}

	return uint(value)
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
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return errValue
	}

	return value
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

// 替换字符串中的 old 为 new，n 为替换的最大次数（小于 0 表示全部替换）
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
func (st StrType) InArray(value string, arr []string) bool {
	return st.Find(value, arr) >= 0
}

// 已排序数组查找，threshold 参数生效
func (st StrType) InSortedArray(value string, arr []string, isAsc bool) bool {
	return st.FindSorted(value, arr, isAsc) >= 0
}

// 遍历查找数组
//
// 如果数组长度较长或对同一数组做多次 LoopFind，建议先 ArrayAsc 后使用 BinFind
//
// 成功时返回查找到的数组下标，失败返回 -1
func (st StrType) LoopFind(value string, arr []string) int {
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
func (st StrType) BinFind(value string, arr []string, isAsc bool) int {
	return st.BinSearch(value, arr, isAsc)
}

// 查找数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找
//
// 二分查找时，将会对数组进行升序排序，查找成功返回的也会是升序后的下标
//
// 成功时返回查找到的数组下标，失败返回 -1
func (st StrType) Find(value string, arr []string) int {
	if len(arr) > threshold {
		return st.SortAndBinSearch(value, arr)
	} else {
		// 否则使用遍历查找
		return st.LoopFind(value, arr)
	}
}

// 查找已排序数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找
//
// 成功时返回查找到的数组下标，失败返回 -1
func (st StrType) FindSorted(value string, arr []string, isAsc bool) int {
	if len(arr) > threshold {
		return st.BinFind(value, arr, isAsc)
	} else {
		// 否则使用遍历查找
		return st.LoopFind(value, arr)
	}
}

func (st StrType) ArrayAsc(arr []string) {
	sort.Strings(arr)
}

func (st StrType) ArrayDesc(arr []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(arr)))
}

// 二分查找(此函数不会进行排序，请输入已排序的数组，否则会产生非预期结果)
// 成功返回查找到的数组下标，失败返回 -1
func (st StrType) BinSearch(value string, arr []string, isAsc bool) int {
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

// 对数组排序并进行二分查找法，成功返回查找到的数组下标，失败返回 -1
func (st StrType) SortAndBinSearch(value string, arr []string) int {
	tmp := make([]string, len(arr))
	copy(tmp, arr)
	st.ArrayAsc(tmp)
	return st.BinSearch(value, tmp, true)
}
