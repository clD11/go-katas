package hackerrank

import "strings"

func HackerrankString(input string) string {
	start := 0
	keys := strings.Split("hackerrank", "")
	for _, key := range keys {
		input = input[start:len(input)]
		start = strings.Index(input, key)
		if start == -1 {
			return "NO"
		}
		start += 1
	}
	return "YES"
}
