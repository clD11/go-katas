package hackerrank

import (
	"github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestStringConstruction_1(t *testing.T) {
	s := "abcd"
	actual := StringConstruction(s)
	testsupport.AssertThatInt32(t, actual, 4)
}

func TestStringConstruction_2(t *testing.T) {
	s := "abab"
	actual := StringConstruction(s)
	testsupport.AssertThatInt32(t, actual, 2)
}
