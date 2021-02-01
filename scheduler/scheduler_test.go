package scheduler_test

import (
	"github.com/clD11/go-katas/scheduler"
	"reflect"
	"testing"
)

func TestShouldSheduleFunctions(t *testing.T) {
	s := scheduler.New()
	s.Schedule(func(i ...int) int {
		t.Fail()
		return 0
	})
}

func TestShouldReturnResultsInOrderScheduled(t *testing.T) {
	s := scheduler.New()
	s.Schedule(func(i ...int) int {
		count := 0
		for _, j := range i {
			count += j
		}
		return count
	})
	s.Schedule(func(i ...int) int {
		return 1
	})
	s.Schedule(func(i ...int) int {
		count := 0
		for _, j := range i {
			count += j
		}
		return count
	}, 5, 5)
	actual := s.Invoke()
	expected := []int{0, 1, 10}
	if !reflect.DeepEqual(actual, expected) {
		t.Fail()
	}
}
