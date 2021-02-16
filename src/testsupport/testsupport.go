package testsupport

import (
	"reflect"
	"testing"
)

func AssertThatInt32(t *testing.T, actual int32, expected int32) {
	if actual != expected {
		t.Errorf("expected %d actual %d", expected, actual)
	}
}

func AssertThatArrInt32(t *testing.T, actual []int32, expected []int32) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %d actual %d", expected, actual)
	}
}
