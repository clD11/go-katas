package hackerrank

import (
	"fmt"
	"strconv"
	"strings"
)

func findBeautifulNumber(input string) string {
	digits := 1
	startIndex := 0
	endIndex := digits

	for {
		num, _ := strconv.Atoi(input[startIndex:endIndex])
		target := strconv.Itoa(num + 1)
		nextValue := strings.Index(input, target)

		if nextValue == endIndex {
			startIndex = nextValue
			endIndex = startIndex + len(target)
		} else {
			digits += 1
			startIndex = 0
			endIndex = digits
		}

		if endIndex >= len(input)-1 {
			if nextValue != -1 {
				return fmt.Sprintf("YES %s", input[0:digits])
			}
			return "NO"
		}

	}
}
