package hackerrank

import "strings"

func GameOfThrones(s string) string {
	freq := make(map[string]int)
	for _, char := range strings.Split(s, "") {
		freq[char]++
	}
	odds := 0
	for _, v := range freq {
		if v%2 != 0 {
			odds++
		}
	}
	if odds > 1 {
		return "NO"
	}
	return "YES"
}
