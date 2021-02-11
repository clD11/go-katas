package hackerrank

func formingMagicSquare(s [][]int32) int32 {
	//// check diagonal
	d1 := makePositive(s[0][0]-5) - makePositive(s[2][2]-5)
	if d1 != 0 {
		s[2][2] = doSwitch(s[0][0])
	}

	m1 := makePositive(s[0][1]-5) - makePositive(s[2][1]-5)
	if m1 != 0 {
		s[2][1] = doSwitch(s[0][1])
	}

	d2 := makePositive(s[0][2]-5) - makePositive(s[2][0]-5)
	if d2 != 0 {
		s[2][0] = doSwitch(s[0][2])
	}

	rm2 := (makePositive(s[1][0]) - 5) - (makePositive(s[1][2]) - 5)
	if rm2 != 0 {
		s[1][2] = doSwitch(s[1][0])
	}

	return 1
}

func doSwitch(n int32) int32 {
	if n > 0 {
	}
	return n
}

func makePositive(n int32) int32 {
	if n < 0 {
		return n * -1
	}
	return n
}
