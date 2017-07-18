package gokatas

func SumDiagonalDifference(arr [][]int) int {
	L := len(arr)
	
	var DL int
	var DR int
	
	for i := 0; i < L; i++ {
		DL = DL + arr[i][i]
		DR = DR + arr[L - i - 1][L - i - 1]
	}

	return DL - DR
}