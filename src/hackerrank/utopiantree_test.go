package hackerrank

import "testing"

func TestReturnHeightForCycle0(t *testing.T) {
	actual := utopianTree(0)
	expected := int32(1)
	assertThat(t, actual, expected)
}

func TestReturnHeightForCycle1(t *testing.T) {
	actual := utopianTree(1)
	expected := int32(2)
	assertThat(t, actual, expected)
}

func TestReturnHeightForCycle2(t *testing.T) {
	actual := utopianTree(2)
	expected := int32(3)
	assertThat(t, actual, expected)
}

func TestReturnHeightForCycle3(t *testing.T) {
	actual := utopianTree(3)
	expected := int32(6)
	assertThat(t, actual, expected)
}

func TestReturnHeightForCycle5(t *testing.T) {
	actual := utopianTree(5)
	expected := int32(14)
	assertThat(t, actual, expected)
}

func assertThat(t *testing.T, actual int32, expected int32) {
	if actual != expected {
		t.Errorf("expected %d actual %d", expected, actual)
	}
}
