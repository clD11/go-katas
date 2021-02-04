package hackerrank

import (
	"testing"
)

func TestDayOfTheProgrammerGregorianLeapYear(t *testing.T) {
	expected := "12.09.2016"
	actual := DayOfTheProgrammer(2016)
	if actual != expected {
		t.Errorf(actual, expected)
	}
}

func TestDayOfTheProgrammerGregorianNonLeapYear(t *testing.T) {
	expected := "13.09.2017"
	actual := DayOfTheProgrammer(2017)
	if actual != expected {
		t.Errorf(actual, expected)
	}
}

func TestDayOfTheProgrammerJulianLeapYear(t *testing.T) {
	expected := "12.09.1800"
	actual := DayOfTheProgrammer(1800)
	if actual != expected {
		t.Errorf(actual, expected)
	}
}

func TestDayOfTheProgrammerRussianTransition(t *testing.T) {
	expected := "26.09.1918"
	actual := DayOfTheProgrammer(1918)
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
		t.Errorf("Should not be a leap year")
	}
}

func TestNotBeLeapYearRussianTransition(t *testing.T) {
	if isLeapYear(1918) == true {
		t.Errorf("Should not be a leap year")
	}
}
