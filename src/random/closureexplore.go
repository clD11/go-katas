package random

func Execute(f func() int32) int32 {
	return f()
}

func Routines(n int) []int {
	result := make([]int, 0, 0)
	done := make(chan bool)
	for i := 0; i < n; i++ {
		go func(done chan bool, index int) {
			result = append(result, index)
			done <- true
		}(done, i)
	}
	for i := 0; i < n; i++ {
		<-done
	}
	return result
}
