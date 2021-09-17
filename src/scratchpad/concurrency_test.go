package scratchpad

import (
	"testing"

	"github.com/clD11/go-katas/src/testsupport"
)

func TestSwapInt32_True(t *testing.T) {
	var i int32 = 10
	actual := CompareAndSwapInt32(&i, 10, 11)
	testsupport.AssertTrue(t, actual)
}

func TestSwapInt32_False(t *testing.T) {
	var i int32 = 10
	actual := CompareAndSwapInt32(&i, 11, 11)
	testsupport.AssertFalse(t, actual)
}
