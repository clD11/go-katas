package gokatas

import "testing"

var arabicRomanPairsTest = []struct {
	in int
	out string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{10, "X"},
	{12, "XII"},
	{40, "XL"},
	{50, "L"},
	{90, "XC"},
	{97, "XCVII"},
	{100, "C"},
	{500, "D"},
	{1000, "M"},
}

func TestConvertArabicToRomanNumeralShouldReturnIWhenGivenOne(t *testing.T) {
	for _, pair := range arabicRomanPairsTest {
		var expected = pair.out
		var actual = ConvertArabicToRomanNumeral(pair.in)
		if expected != actual {
			t.Error("Test Failed with args ", pair.in, ": expected ", expected, " actual ", actual)
		}
	}
}