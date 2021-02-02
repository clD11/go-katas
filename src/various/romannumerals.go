package various

var arabicRomanPairs = []struct {
	in  int
	out string
}{
	{1000, "M"},
	{500, "D"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertArabicToRomanNumeral(arabic int) string {
	roman := ""
	for _, pair := range arabicRomanPairs {
		for arabic >= pair.in {
			roman += pair.out
			arabic -= pair.in
		}

		if arabic < 1 {
			break
		}
	}

	return roman
}
