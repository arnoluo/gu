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
