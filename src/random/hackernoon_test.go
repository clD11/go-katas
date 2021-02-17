package random

import (
	. "github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestFindMissingNumber(t *testing.T) {
	input := []int32{1, 3, 6, 7, 8, 9, 4, 5, 10}
	actual := findMissingNumber(input)
	AssertThatInt32(t, actual, 2)
}

func TestNotFindMissingNumber(t *testing.T) {
	input := []int32{1, 2, 3, 4, 6, 7, 8, 9, 5, 10}
	actual := findMissingNumber(input)
	AssertThatInt32(t, actual, -1)
}
