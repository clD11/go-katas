package hackerrank

import "strings"

func StringConstruction(s string) int32 {
	cost := 0
	chars := make(map[string]int)
	for _, char := range strings.Split(s, "") {
		_, ok := chars[char]
		if !ok {
			chars[char]++
			cost++
		}
	}
	return int32(cost)
}
