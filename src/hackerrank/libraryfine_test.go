package hackerrank

import "testing"

func TestCalculateFineDays(t *testing.T) {
	actual := libraryFine(9, 6, 2015, 6, 6, 2015)
	expected := int32(45)
	assertThat(t, actual, expected)
}

func TestCalculateFineDays2(t *testing.T) {
	actual := libraryFine(15, 7, 2014, 1, 7, 2015)
	expected := int32(0)
	assertThat(t, actual, expected)
}

func TestCalculateFineMonths(t *testing.T) {
	actual := libraryFine(6, 8, 2015, 6, 6, 2015)
	expected := int32(1000)
	assertThat(t, actual, expected)
}

func TestCalculateFineYears(t *testing.T) {
	actual := libraryFine(6, 8, 2015, 6, 6, 2014)
	expected := int32(10000)
	assertThat(t, actual, expected)
}

func TestCalculateNoFine(t *testing.T) {
	actual := libraryFine(5, 6, 2015, 6, 6, 2015)
	expected := int32(0)
	assertThat(t, actual, expected)
}
