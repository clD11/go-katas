package hackerrank

import "testing"

func TestReturnHeightForCycle0(t *testing.T) {
	actual := utopianTree(0)
	expected := int32(1)
	if actual != expected {
		t.Errorf("expected %d actual %d", expected, actual)
	}
}

func TestReturnHeightForCycle1(t *testing.T) {
	actual := utopianTree(1)
	expected := int32(2)
	if actual != expected {
		t.Errorf("expected %d actual %d", expected, actual)
	}
}

func TestReturnHeightForCycle2(t *testing.T) {
	actual := utopianTree(2)
	expected := int32(3)
	if actual != expected {
		t.Errorf("expected %d actual %d", expected, actual)
	}
}

func TestReturnHeightForCycle3(t *testing.T) {
	actual := utopianTree(3)
	expected := int32(6)
	if actual != expected {
		t.Errorf("expected %d actual %d", expected, actual)
	}
}

func TestReturnHeightForCycle5(t *testing.T) {
	actual := utopianTree(5)
	expected := int32(14)
	if actual != expected {
		t.Errorf("expected %d actual %d", expected, actual)
	}
}
