package hackerrank

import "strings"

func twoStrings(s1 string, s2 string) string {
	var longest string
	var shortest string
	if len(s1) > len(s2) {
		longest = s1
		shortest = s2
	} else {
		longest = s2
		shortest = s1
	}
	for _, letter := range strings.Split(shortest, "") {
		if strings.Contains(longest, letter) {
			return "YES"
		}
	}
	return "NO"
}
