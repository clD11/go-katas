package random

import (
	"github.com/clD11/go-katas/src/testsupport"
	"reflect"
	"sort"
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

func TestRoutines(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4}
	actual := Routines(len(expected))
	sort.Ints(actual)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v actual: %v", expected, actual)
	}
}
