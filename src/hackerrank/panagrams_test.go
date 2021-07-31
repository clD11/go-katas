package hackerrank

import (
	"github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestPangram_Pangram(t *testing.T) {
	input := "We promptly judged antique ivory buckles for the next prize"
	actual := Pangram(input)
	testsupport.AssertThatString(t, actual, "pangram")
}

func TestPangram_NotPangram(t *testing.T) {
	input := "We promptly judged antique ivory buckles for the prize"
	actual := Pangram(input)
	testsupport.AssertThatString(t, actual, "not pangram")
}
