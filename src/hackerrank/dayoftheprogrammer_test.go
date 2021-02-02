package hackerrank

import (
	"testing"
)

func TestDayOfTheProgrammer(t *testing.T) {
	expected := "12.09.2016"
	actual := DayOfTheProgrammer(2016)
	if actual != expected {
		t.Errorf(actual, expected)
	}
}

func TestIsLeapYearGregorian(t *testing.T) {
	if isLeapYear(1984) == false {
		t.Errorf("Should be a leap year")
	}
}

func TestIsLeapYearJulian(t *testing.T) {
	if isLeapYear(1800) == false {
		t.Errorf("Should be a leap year")
	}
}

func TestNotBeLeapYear(t *testing.T) {
	if isLeapYear(2019) == true {
		t.Errorf("Should be a leap year")
	}
}
