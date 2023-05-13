package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var st StrType

func TestStrIf(t *testing.T) {
	assert.Equal(t, "trueValue", st.If(true, "trueValue", "falseValue"), "should be string: trueValue")
	assert.Equal(t, "falseValue", st.If(false, "trueValue", "falseValue"), "should be string: falseValue")
}

func TestSub(t *testing.T) {
	assert.Equal(t, st.Sub("abcdef", 0, 2), "ab")
	assert.Equal(t, st.Sub("测试中", 0, 2), "测试")
}

func TestRandom(t *testing.T) {
	a := st.Rand(6)
	b := st.RandLetters(6)
	c := st.RandNumbers(6)
	d := st.RandChars("c", 3)
	fmt.Println(a, b, c, d)
	assert.Len(t, a, 6)
	assert.Len(t, b, 6)
	assert.Len(t, c, 6)
	assert.Equal(t, "ccc", d)
}

func TestRegReplace(t *testing.T) {
	c := st.RegReplace("acacac", `a+`, "c")
	assert.Equal(t, "cccccc", c)
}

func TestConvert(t *testing.T) {
	var a, b, c, d string = "1", "-1", "a", "0"
	assert.Equal(t, true, st.Bool(a, false))
	assert.Equal(t, 1, st.Int(a, -1))
	assert.Equal(t, uint(1), st.Pint(b, 1))
	assert.Equal(t, uint(0), st.Uint(d, 1))
	assert.Equal(t, false, st.IsEmpty(c))
	assert.Equal(t, false, st.IsInt(c))
	assert.Equal(t, true, st.IsInt(a))
	assert.Equal(t, true, st.IsNum(b))
	assert.Equal(t, false, st.IsNum(c))

}

func TestStringFind(t *testing.T) {
	var arr = []string{"b", "d", "e", "c", "f", "i", "h", "g", "a", "j"}
	var arrAsc = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	var arrDesc = []string{"j", "i", "h", "g", "f", "e", "d", "c", "b", "a"}

	item := "d"
	pos := st.Find(item, arr)
	// 二分查找，返回的是升序后的下标
	assert.Equal(t, 3, pos)

	pos = st.FindSorted(item, arrAsc, true)
	assert.Equal(t, 3, pos)

	pos = st.FindSorted(item, arrDesc, false)
	assert.Equal(t, 6, pos)

	pos = st.LoopFind(item, arr)
	assert.Equal(t, 1, pos)

	pos = st.BinFind(item, arrAsc, true)
	assert.Equal(t, 3, pos)

	pos = st.BinFind(item, arrDesc, false)
	assert.Equal(t, 6, pos)

	pos = st.SortAndBinSearch(item, arr)
	assert.Equal(t, 3, pos)

	fmt.Println(arr)
}

func TestStrUtils(t *testing.T) {
	assert.Equal(t, 5.5, st.Float("5.5", 0.0))
	assert.Equal(t, 0.0, st.Float("aa5.5", 0.0))
	s := "abcdefg"
	assert.Equal(t, 7, st.Len(s))
	assert.Equal(t, "Abcdefg", st.UpperFirst(s))
	assert.Equal(t, "abc", st.LowerFirst("Abc"))
	assert.Equal(t, "ab", st.TrimSpace(" ab "))
	assert.Equal(t, "ab", st.Trim("-ab-", "-"))
	assert.Equal(t, "ab-", st.Ltrim("-ab-", "-"))
	assert.Equal(t, "-ab", st.Rtrim("-ab-", "-"))
	assert.Equal(t, "a|b", st.Join([]string{"a", "b"}, "|"))
	assert.Equal(t, []string{"a", "b"}, st.Split("a|b", "|"))
	assert.Equal(t, -1, st.Index("abc", "d"))
	assert.Equal(t, 0, st.Index("abc", "a"))
	assert.Equal(t, "bbbbb", st.Replace("ababa", "a", "b", -1))
	assert.Equal(t, "baba", st.Replace("ababa", "a", "", 1))
	assert.True(t, st.HasPrefix("abc", "a"))
	assert.True(t, st.HasSuffix("abc", "c"))
	assert.Equal(t, "abcde", st.Format("abc%s", "de"))

	inArr := st.InArray(s, []string{"abc", "abcde", "abcdef", "abcdefg"})
	assert.True(t, inArr)

	inArr2 := st.InSortedArray(s, []string{"abc", "abcde", "abcdef", "abcdefg"}, true)
	assert.True(t, inArr2)
}

func TestStringSort(t *testing.T) {
	var arr = []string{"a", "c", "d", "b", "e"}
	var arrAsc = []string{"a", "b", "c", "d", "e"}
	var arrDesc = []string{"e", "d", "c", "b", "a"}

	st.ArrayAsc(arr)
	assert.Equal(t, arrAsc, arr)

	st.ArrayDesc(arr)
	assert.Equal(t, arrDesc, arr)
}
