package hackerrank

import (
	. "github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestEqualizeTheArrayA(t *testing.T) {
	arr := []int32{1, 2, 2, 3}
	expected := int32(2)
	actual := equalizeArray(arr)
	AssertThatInt32(t, actual, expected)
}

func TestEqualizeTheArrayB(t *testing.T) {
	arr := []int32{10, 27, 9, 10, 100, 38, 30, 32, 45, 29, 27, 29, 32, 38, 32, 38, 14, 38, 29, 30, 63, 29, 63, 91, 54, 10, 63}
	expected := int32(23)
	actual := equalizeArray(arr)
	AssertThatInt32(t, actual, expected)
}
