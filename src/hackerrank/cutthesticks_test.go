package hackerrank

import "testing"

func TestCutSticksA(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 3, 3, 2, 1}
	expected := []int32{8, 6, 4, 1}
	actual := cutTheSticks(arr)
	assertThatArrInt32(t, actual, expected)
}

func TestCutSticksB(t *testing.T) {
	arr := []int32{1, 2, 3}
	expected := []int32{3, 2, 1}
	actual := cutTheSticks(arr)
	assertThatArrInt32(t, actual, expected)
}

func TestCutSticksC(t *testing.T) {
	arr := []int32{5, 4, 4, 2, 2, 8}
	expected := []int32{6, 4, 2, 1}
	actual := cutTheSticks(arr)
	assertThatArrInt32(t, actual, expected)
}

func TestCutSticksD(t *testing.T) {
	arr := []int32{2, 6, 8, 9}
	expected := []int32{4, 3, 2, 1}
	actual := cutTheSticks(arr)
	assertThatArrInt32(t, actual, expected)
}
