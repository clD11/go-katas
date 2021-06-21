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

func TestFindBeautiful_Mix(t *testing.T) {
	input := "7891011"
	expected := "YES 7"
	actual := findBeautifulNumber(input)
	assertBeauty(t, actual, expected)
}

func TestFindBeautiful(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"99910001001", "YES 999"},
		{"7891011", "YES 7"},
		{"9899100", "YES 98"},
		{"999100010001", "NO"},
	}
	for _, e := range tests {
		actual := findBeautifulNumber(e.input)
		assertBeauty(t, actual, e.expected)
	}
}

func TestFindBeautifulLong(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"57585960616263646566676869707172", "YES 57"},
		{"57585960616263646566676869707072", "NO"},
		{"57585960616263646566676869707172", "YES 57"},
		{"57585960616263646566676869706172", "NO"},
		{"89101112131415161718192021222324", "YES 8"},
		{"89101112131415161718192021212324", "NO"},
		{"42434445464748495051525354555657", "YES 42"},
		{"42434445464748495051525354455657", "NO"},
		{"858687888990919293949596979899", "YES 85"},
		{"858687888990919293949595979899", "NO"},
	}
	for _, e := range tests {
		actual := findBeautifulNumber(e.input)
		assertBeauty(t, actual, e.expected)
	}
}

func assertBeauty(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("expected %s actual %s", expected, actual)
	}
}
