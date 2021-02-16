package hackerrank

import (
	. "github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestReturnHeightForCycle0(t *testing.T) {
	actual := utopianTree(0)
	expected := int32(1)
	AssertThatInt32(t, actual, expected)
}

func TestReturnHeightForCycle1(t *testing.T) {
	actual := utopianTree(1)
	expected := int32(2)
	AssertThatInt32(t, actual, expected)
}

func TestReturnHeightForCycle2(t *testing.T) {
	actual := utopianTree(2)
	expected := int32(3)
	AssertThatInt32(t, actual, expected)
}

func TestReturnHeightForCycle3(t *testing.T) {
	actual := utopianTree(3)
	expected := int32(6)
	AssertThatInt32(t, actual, expected)
}

func TestReturnHeightForCycle5(t *testing.T) {
	actual := utopianTree(5)
	expected := int32(14)
	AssertThatInt32(t, actual, expected)
}
