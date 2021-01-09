package hackerrank

import (
	"sort"
	"strconv"
)

func CalculateMinMaxSum(arg []float64) string {
	sort.Float64s(arg)

	var total float64
	for _, num := range arg {
		total = total + num
	}

	l := strconv.FormatFloat((total - arg[len(arg)-1]), 'f', -1, 64)
	h := strconv.FormatFloat((total - arg[0]), 'f', -1, 64)

	return l + " " + h
}
