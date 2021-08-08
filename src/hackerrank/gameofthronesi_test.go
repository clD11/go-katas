package hackerrank

import (
	"github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestGameOfThrones_YES(t *testing.T) {
	s := "aaabbbb"
	actual := GameOfThrones(s)
	testsupport.AssertThatString(t, actual, "YES")
}

func TestGameOfThrones_NO(t *testing.T) {
	s := "cdefghmnopqrstuvw"
	actual := GameOfThrones(s)
	testsupport.AssertThatString(t, actual, "NO")
}

func TestGameOfThrones_Longer_NO(t *testing.T) {
	s := "cdcdcdcdeeeef"
	actual := GameOfThrones(s)
	testsupport.AssertThatString(t, actual, "YES")
}
