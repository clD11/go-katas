package hackerrank

func anagram(s string) int32 {
	if len(s)%2 != 0 {
		return -1
	}

	freq := make(map[string]int)

	for i := 0; i < len(s)/2; i++ {
		freq[string(s[i])]++
		freq[string(s[len(s)-1-i])]--
	}

	count := 0
	for _, v := range freq {
		if v > 0 {
			count += v
		}
	}

	return int32(count)
}
