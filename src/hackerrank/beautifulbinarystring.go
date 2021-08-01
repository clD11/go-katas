package hackerrank

import "strings"

const bbs = "010"

func BeautifulBinaryString(input string) int {
	count := 0
	index := 0
	for {
		index = strings.Index(input, bbs)
		if index == -1 {
			break
		}
		input = input[index+3:]
		count += 1
	}
	return count
}
