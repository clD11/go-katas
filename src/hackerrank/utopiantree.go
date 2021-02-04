package hackerrank

func utopianTree(n int32) int32 {
	height := int32(1)
	for i := 1; i <= int(n); i++ {
		if height == 1 {
			height += 1
		} else if height%2 == 0 {
			height += 1
		} else {
			height *= 2
		}
	}
	return height
}
