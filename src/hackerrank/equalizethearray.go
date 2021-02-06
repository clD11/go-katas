package hackerrank

func equalizeArray(arr []int32) int32 {
	frequency := make([]int32, 101, 101)
	for i := 0; i < len(arr); i++ {
		frequency[arr[i]]++
	}

	index := 0
	freq := int32(0)
	for i := 0; i < len(frequency); i++ {
		if frequency[i] > freq {
			freq = frequency[i]
			index = i
		}
	}

	equalize := int32(0)
	for i := 0; i < len(frequency); i++ {
		if i != index {
			equalize += frequency[i]
		}
	}

	return equalize
}
