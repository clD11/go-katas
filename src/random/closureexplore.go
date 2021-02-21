package random

func Execute(f func() int32) int32 {
	return f()
}
