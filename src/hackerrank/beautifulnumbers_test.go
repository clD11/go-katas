package hackerrank

import "testing"

func TestFindBeautiful_Nines(t *testing.T) {
	input := "198199200"
	expected := "YES 198"
	actual := findBeautifulNumber(input)
	assertBeauty(t, actual, expected)
}

func TestFindBeautiful_MoreNines(t *testing.T) {
	input := "99100"
	expected := "YES 99"
	actual := findBeautifulNumber(input)
	assertBeauty(t, actual, expected)
}

func TestFindBeautiful_Consecutive(t *testing.T) {
	input := "1234"
	expected := "YES 1"
	actual := findBeautifulNumber(input)
	assertBeauty(t, actual, expected)
}

func TestFindBeautiful_nine(t *testing.T) {
	input := "91011"
	expected := "YES 9"
	actual := findBeautifulNumber(input)
	assertBeauty(t, actual, expected)
}

func TestFindBeautiful_No(t *testing.T) {
	input := "101103"
	expected := "NO"
	actual := findBeautifulNumber(input)
	assertBeauty(t, actual, expected)
}

func TestFindBeautiful_NoZeros(t *testing.T) {
	input := "010203"
	expected := "NO"
	actual := findBeautifulNumber(input)
	assertBeauty(t, actual, expected)
}

func assertBeauty(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("expected %s actual %s", expected, actual)
	}
}
