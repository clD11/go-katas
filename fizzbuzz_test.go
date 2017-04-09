package gokatas

import (
	"testing"
)

var testData = []struct {
	in int
	out string
}{
	{1,"1"},
	{2,"2"},
	{3,"Fizz"},
	{5,"Buzz"},
	{15, "FizzBuzz"},
	{25, "Buzz"},
	{30, "FizzBuzz"},
	{33, "Fizz"},
	{100, "Buzz"},
}

func TestFizzBuzzShouldReturnFizzBuzz(t *testing.T) {
	for _, pair := range testData {
		expected := pair.out
		actual := FizzBuzz(pair.in)
		if expected != actual {
			t.Error("Test failed expected ", expected, " actual ", actual)
		}
	}
}




