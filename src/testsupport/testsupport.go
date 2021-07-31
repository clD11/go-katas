package testsupport

import (
	"reflect"
	"testing"
)

func AssertThatInt(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Errorf("expected %d actual %d", expected, actual)
	}
}

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

func AssertThatString(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("expected %s actual %s", expected, actual)
	}
}
