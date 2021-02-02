package hackerrank

import (
	"strconv"
	"strings"
)

func Sum(input string) int64 {
	summable := strings.Split(input, " ")

	var sum int64
	for _, num := range summable {
		if n, err := strconv.ParseInt(num, 10, 64); err == nil {
			sum = sum + n
		}
	}

	return sum
}
