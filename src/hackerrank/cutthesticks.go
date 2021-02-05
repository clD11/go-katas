package hackerrank

func cutTheSticks(arr []int32) []int32 {
	candles := make([]int32, 1001, 1001)
	for i := 0; i < len(arr); i++ {
		candles[arr[i]]++
	}
	cuts := make([]int32, 0, 0)
	totalCandles := int32(len(arr))
	for i := 1; i < len(candles); i++ {
		if candles[i] > 0 {
			cuts = append(cuts, totalCandles)
			totalCandles -= candles[i]
		}
	}
	return cuts
}
