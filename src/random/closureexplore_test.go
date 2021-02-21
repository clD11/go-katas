package random

import (
	"github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestClosure(t *testing.T) {
	var i int32 = 1
	f := func() int32 {
		return i + i
	}
	actual := Execute(f)
	testsupport.AssertThatInt32(t, actual, 2)
}
