package hackerrank

import (
	"github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestLoveLetterMystery_1(t *testing.T) {
	s := "abc"
	actual := LoveLetterMystery(s)
	testsupport.AssertThatInt32(t, actual, 2)
}

func TestLoveLetterMystery_2(t *testing.T) {
	s := "abcba"
	actual := LoveLetterMystery(s)
	testsupport.AssertThatInt32(t, actual, 0)
}

func TestLoveLetterMystery_3(t *testing.T) {
	s := "abcd"
	actual := LoveLetterMystery(s)
	testsupport.AssertThatInt32(t, actual, 4)
}

func TestLoveLetterMystery_4(t *testing.T) {
	s := "cba"
	actual := LoveLetterMystery(s)
	testsupport.AssertThatInt32(t, actual, 2)
}
