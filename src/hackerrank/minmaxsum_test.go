package hackerrank

import (
	"testing"
)

func TestCalculateMinMaxSum(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}

	expected := "10 14"
	actual := CalculateMinMaxSum(input)

	if expected != actual {
		t.Error("expected", expected, "actual", actual)
	}
}
