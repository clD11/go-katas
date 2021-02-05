package hackerrank

import "testing"

func assertThat(t *testing.T, actual int32, expected int32) {
	if actual != expected {
		t.Errorf("expected %d actual %d", expected, actual)
	}
}
