package scratchpad

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	p := Pointer{}

	p2 := Pointer{
		createdAt: p.createdAt,
	}

	if p2.createdAt != nil {
		fmt.Errorf("bug")
	}
}

func TestNilPointerDereference(t *testing.T) {
	s := make([]Data, 0, 1)

	r := &Data{
		ID: "nil",
	}

	s = append(s, *r)

	if s[0].ID != "nil" {
		fmt.Errorf("bug")
	}
}

func TestFindMatchesAndCombine(t *testing.T) {
	requested := makeRefunds(10)

	r1 := Data{
		ID:    "a",
		State: "requested",
	}

	r2 := Data{
		ID:    "b",
		State: "requested",
	}

	requested = append(requested, r1, r2)

	s1 := Data{
		ID:    "a",
		State: "success",
	}

	s2 := Data{
		ID:    "b",
		State: "success",
	}

	success := make([]Data, 0, 2)
	success = append(success, s1, s2)

	actual := FindMatchesAndCombine(success, requested)
	ops := FindMatchesAndCombineOp(success, requested)

	reflect.DeepEqual(actual, ops)
}

func makeRefunds(n int) []Data {
	r := make([]Data, n)
	for i := 0; i < n; i++ {
		r[i] = randomRefund()
	}
	return r
}

func randomRefund() Data {
	return Data{
		ID:    time.Now().String(),
		State: "requested",
	}
}
