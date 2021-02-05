package hackerrank

import (
	"reflect"
	"testing"
)

func assertThatInt32(t *testing.T, actual int32, expected int32) {
	if actual != expected {
		t.Errorf("expected %d actual %d", expected, actual)
	}
}

func assertThatArrInt32(t *testing.T, actual []int32, expected []int32) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %d actual %d", expected, actual)
	}
}
