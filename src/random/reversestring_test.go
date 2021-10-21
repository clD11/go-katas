package random

import (
	"github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestReverseString_Zero(t *testing.T) {
	input := "dog"
	expected := "dog"
	actual := ReverseString(input)
	testsupport.AssertThatString(t, actual, expected)
}

func TestReverseString_One(t *testing.T) {
	input := "big dog"
	expected := "dog big"
	actual := ReverseString(input)
	testsupport.AssertThatString(t, actual, expected)
}

func TestReverseString_Two(t *testing.T) {
	input := "along came polly"
	expected := "polly came along"
	actual := ReverseString(input)
	testsupport.AssertThatString(t, actual, expected)
}
