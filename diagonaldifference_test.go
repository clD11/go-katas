package gokatas

import (
	"testing"
)

func testSumDiagonalDifference(t *testing.T) {
	input := [][]int{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}

	expected := 15
	actual := SumDiagonalDifference(input)

	if actual != expected {
		t.Error("Error expected ", expected, "actual", actual)
	}
}
