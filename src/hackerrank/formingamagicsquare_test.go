package hackerrank

import (
	"fmt"
	"testing"
)

func TestFormMagicSquare(t *testing.T) {
	s := [][]int32{{4, 9, 2}, {3, 5, 7}, {8, 1, 5}}
	expected := int32(1)
	actual := formingMagicSquare(s)
	assertThatInt32(t, actual, expected)
}

func TestNegativePositive(t *testing.T) {
	fmt.Printf("%d ", -1*-1)
}
