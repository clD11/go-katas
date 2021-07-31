package hackerrank

import (
	"github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestTwoStrings_YES(t *testing.T) {
	s1 := "hello"
	s2 := "world"
	actual := twoStrings(s1, s2)
	testsupport.AssertThatString(t, actual, "YES")
}

func TestTwoStrings_NO(t *testing.T) {
	s1 := "hi"
	s2 := "world"
	actual := twoStrings(s1, s2)
	testsupport.AssertThatString(t, actual, "NO")
}
