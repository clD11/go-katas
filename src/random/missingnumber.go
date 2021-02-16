package random

import (
	"sort"
)

// Could use n(n+1)/2 but wheres the fun in that

func find(input []int32) int32 {
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	for i := 0; i < len(input)-1; i++ {
		if input[i+1]-input[i] != 1 {
			return input[i] + 1
		}
	}
	return -1
}
