package random

import (
	"reflect"
	"testing"
)

func TestGetPlayerFrequencyTopTwo(t *testing.T) {
	actual := GetPlayerFrequency("testdata/players.txt", 2)
	expected := []string{"Anna", "Joe"}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(actual, "is not", expected)
	}
}

func TestGetPlayerFrequencyTopThree(t *testing.T) {
	actual := GetPlayerFrequency("testdata/players.txt", 3)
	expected := []string{"Anna", "Joe", "Bob"}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(actual, "is not", expected)
	}
}

func TestFrequency(t *testing.T) {
	input := setupBlog()
	expected := []string{"art", "business", "entertainment"}
	actual := Frequency(input, 3)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("actual: %s expected: %s", actual, expected)
	}
}

func setupBlog() []Blog {
	b := make([]Blog, 0, 0)

	title := "art"
	for i := 0; i < 5; i++ {
		b = append(b, Blog{
			title: title,
			like:  true,
		})
	}

	title = "business"
	for i := 0; i < 4; i++ {
		b = append(b, Blog{
			title: title,
			like:  true,
		})
	}

	title = "entertainment"
	for i := 0; i < 3; i++ {
		b = append(b, Blog{
			title: title,
			like:  true,
		})
	}

	title = "sport"
	for i := 0; i < 2; i++ {
		b = append(b, Blog{
			title: title,
			like:  true,
		})
	}

	return b
}
