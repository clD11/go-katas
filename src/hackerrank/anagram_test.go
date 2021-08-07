package hackerrank

import (
	"github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestAnagram_1(t *testing.T) {
	s := "aaabbb"
	actual := anagram(s)
	testsupport.AssertThatInt32(t, actual, 3)
}

func TestAnagram_2(t *testing.T) {
	s := "ab"
	actual := anagram(s)
	testsupport.AssertThatInt32(t, actual, 1)
}

func TestAnagram_NotEqual(t *testing.T) {
	s := "abc"
	actual := anagram(s)
	testsupport.AssertThatInt32(t, actual, -1)
}

func TestAnagram_3(t *testing.T) {
	s := "mnop"
	actual := anagram(s)
	testsupport.AssertThatInt32(t, actual, 2)
}

func TestAnagram_AlreadyAnagram(t *testing.T) {
	s := "xyyx"
	actual := anagram(s)
	testsupport.AssertThatInt32(t, actual, 0)
}

func TestAnagram_5(t *testing.T) {
	s := "xaxbbbxx"
	actual := anagram(s)
	testsupport.AssertThatInt32(t, actual, 1)
}
