package glob4go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleMatch(t *testing.T) {
	pattern := []byte("foo*bar*quuk")
	s := []byte("fooulakdjbarlakdjquuk")
	assert.True(t, Glob(pattern, s, false), "glob pattern should match")
}

func TestGroupMatch(t *testing.T) {
	pattern := []byte("abcd[0-9]e*")
	s1 := []byte("abcd7e")
	s2 := []byte("abcd2e898")
	s3 := []byte("abcde")
	assert.True(t, Glob(pattern, s1, false), "match should match")
	assert.True(t, Glob(pattern, s2, false), "match should match")
	assert.False(t, Glob(pattern, s3, false), "match should not match")
}

func TestCaseMatch(t *testing.T) {
	pattern := []byte("ABCDe*")
	s1 := []byte("AbcdE")
	s2 := []byte("ABCDeuuu")
	assert.False(t, Glob(pattern, s1, false), "s1 should not match because of case")
	assert.True(t, Glob(pattern, s2, false), "s2 should match case-sensitively")
}

func TestNegativeMatch(t *testing.T) {
	pattern := []byte("abcd[^xyz]")
	s1 := []byte("abcdx")
	s2 := []byte("abcdy")
	s3 := []byte("abcdz")
	s4 := []byte("abcdu")
	s5 := []byte("abcdux")
	assert.False(t, Glob(pattern, s1, false), "should not match because negative group")
	assert.False(t, Glob(pattern, s2, false), "should not match because negative group")
	assert.False(t, Glob(pattern, s3, false), "should not match because negative group")
	assert.True(t, Glob(pattern, s4, false), "should match because s4 contains no char in negative group")
	assert.False(t, Glob(pattern, s5, false), "should not match because s5 contains extra char at the end")
}

func TestMatchAll(t *testing.T) {
	pattern := []byte("*")
	s1 := []byte("alkdjalkdfj")
	s2 := []byte("ewiuiu")
	assert.True(t, Glob(pattern, s1, false), "* should match all")
	assert.True(t, Glob(pattern, s2, false), "* should match all")
}

func TestSideEffectFreeness(t *testing.T) {
	pattern := []byte("abc*f?")
	s := []byte("abcuidaf8")
	pLen := len(pattern)
	sLen := len(s)
	match := Glob(pattern, s, false)
	assert.True(t, match, "pattern is a match")
	assert.Equal(t, pLen, len(pattern), "the glob() func must not change input slice")
	assert.Equal(t, sLen, len(s), "the glob() func must not change input slice")
}
