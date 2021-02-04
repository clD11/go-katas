package hackerrank

func SumDiagonalDifference(arr [][]int) int {
	L := len(arr)

	var DL int
	var DR int

	for i := 0; i < L; i++ {
		DL = DL + arr[i][i]
		DR = DR + arr[i][L-i-1]
	}

	return DL - DR
}