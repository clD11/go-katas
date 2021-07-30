package random

import (
	"github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestReverseString(t *testing.T) {
	input := "big dog"
	expected := "dog big"
	actual := ReverseString(input)
	testsupport.AssertThatString(t, actual, expected)
}
