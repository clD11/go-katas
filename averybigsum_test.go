package gokatas

import (
	"testing"
)

func testSum(t *testing.T) {
	input := "1000000001 1000000002 1000000003 1000000004 1000000005"

	var expected int64 = 5000000015
	actual := Sum(input)

	if actual != expected {
		t.Error("Error expected ", expected, "actual", actual)
	}
}
