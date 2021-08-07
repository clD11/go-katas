package hackerrank

func PalindromeIndex(input string) int32 {
	posLeft := 0
	posRight := len(input) - 1
	l := -1
	r := -1
	for {
		if posLeft >= posRight {
			break
		}
		if input[posLeft] == input[posRight] {
			posLeft++
			posRight--
		} else {
			if l == -1 {
				l = posLeft
				r = posRight
				posLeft++
			} else {
				return int32(r)
			}
		}
	}
	return int32(l)
}
