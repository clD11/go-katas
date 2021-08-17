package hackerrank

// Alternative approach could use a map to find all possible weights and lookup
func WeightedUniformStrings(s string, queries []int32) []string {
	w := make([]int32, 26, 26)

	target := 0
	var count int32 = 0

	for i := 0; i < len(s); i++ {
		if s[i] == s[target] {
			count++
		} else {
			target = i
			// found new target so start at 1
			count = 1
		}
		if w[s[target]-97] < count {
			w[s[target]-97] = count
		}
	}

	r := make([]string, len(queries), len(queries))

	for index, query := range queries {
		for j := 1; j <= 26; j++ {
			isFactor := query%int32(j) == 0
			if isFactor {
				max := query / int32(j)
				if max <= w[j-1] {
					r[index] = "Yes"
					break
				}
			}
		}
		if r[index] == "" {
			r[index] = "No"
		}
	}
	return r
}
