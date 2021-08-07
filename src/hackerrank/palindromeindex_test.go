package hackerrank

import (
	"github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestPalindromeIndex_End(t *testing.T) {
	input := "aaab"
	actual := PalindromeIndex(input)
	testsupport.AssertThatInt32(t, actual, 3)
}

func TestPalindromeIndex_None(t *testing.T) {
	input := "aaa"
	actual := PalindromeIndex(input)
	testsupport.AssertThatInt32(t, actual, -1)
}

func TestPalindromeIndex_Start(t *testing.T) {
	input := "baaa"
	actual := PalindromeIndex(input)
	testsupport.AssertThatInt32(t, actual, 0)
}

func TestPalindromeIndex_Middle(t *testing.T) {
	input := "cadabadac"
	actual := PalindromeIndex(input)
	testsupport.AssertThatInt32(t, actual, -1)
}

func TestPalindromeIndex_Right_Middle(t *testing.T) {
	input := "caabbadac"
	actual := PalindromeIndex(input)
	testsupport.AssertThatInt32(t, actual, 6)
}

func TestPalindromeLong(t *testing.T) {
	input := "hgygsvlfcwnswtuhmyaljkqlqjjqlqkjlaymhutwsnwcwflvsgygh"
	actual := PalindromeIndex(input)
	testsupport.AssertThatInt32(t, actual, 44)
}

func TestPalindromeLeft(t *testing.T) {
	input := "hgygsvlfwcwnswtuhmyaljkqlqjjqlqkjlaymhutwsnwcflvsgygh"
	actual := PalindromeIndex(input)
	testsupport.AssertThatInt32(t, actual, 8)
}
