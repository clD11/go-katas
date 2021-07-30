package random

import (
	"strings"
)

func ReverseString(input string) string {
	result := ""
	arr := strings.Split(input, " ")
	for i := len(arr) - 1; i >= 0; i-- {
		result += arr[i]
		result += " "
	}
	return result[0 : len(result)-1]
}
