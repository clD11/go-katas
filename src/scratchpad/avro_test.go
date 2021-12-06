package scratchpad

import (
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	expected := SimpleRecord{
		F1:           1,
		F2:           "tester2",
		F3:           "tester3",
		Dependencies: []string{"tester4", "tester5"},
	}

	actual := encodeDecode(expected)

	if !reflect.DeepEqual(actual, expected) {
		t.Error("should be equal")
	}
}
