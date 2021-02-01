package various

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
