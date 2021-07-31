package hackerrank

func Pangram(input string) string {
	upper := make([]int8, 26)
	lower := make([]int8, 26)

	for i := 0; i < len(input); i++ {
		char := input[i]
		if char >= 65 && char <= 90 {
			upper[char-65]++
		}
		if char >= 97 && char <= 122 {
			lower[char-97]++
		}
	}

	for i := 0; i < 26; i++ {
		if upper[i] == 0 && lower[i] == 0 {
			return "not pangram"
		}
	}

	return "pangram"
}
