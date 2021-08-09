package hackerrank

func LoveLetterMystery(s string) int32 {
	i := 0
	j := len(s) - 1

	var reductions int32 = 0

	for {
		if i >= j {
			break
		}
		if s[i] > s[j] {
			reductions += int32(s[i] - s[j])
		} else {
			if s[i] < s[j] {
				reductions += int32(s[j] - s[i])
			}
		}
		i++
		j--
	}

	return reductions
}
