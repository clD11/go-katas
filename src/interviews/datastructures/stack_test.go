package datastructures

import (
	"github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestStackPush(t *testing.T) {
	s := newStack()
	n := &Node{
		value: 1,
	}
	s.Push(n)

	m := &Node{
		value: 2,
	}

	s.Push(m)

	testsupport.AssertThatInt(t, s.values[0].value, n.value)
	testsupport.AssertThatInt(t, s.values[1].value, m.value)
}

func TestStackPop(t *testing.T) {
	s := newStack()
	n := &Node{
		value: 1,
	}
	s.Push(n)

	actual := s.Pop()

	testsupport.AssertThatInt(t, actual.value, n.value)
	testsupport.AssertThatInt(t, s.Len(), 0)
}

func TestPushPop(t *testing.T) {
	s := newStack()

	n := &Node{
		value: 1,
	}
	m := &Node{
		value: 2,
	}
	k := &Node{
		value: 2,
	}

	s.Push(n)
	s.Push(m)
	s.Push(k)

	testsupport.AssertThatInt(t, 3, s.Len())
	testsupport.AssertThatInt(t, s.Pop().value, k.value)
	testsupport.AssertThatInt(t, s.Pop().value, m.value)
	testsupport.AssertThatInt(t, s.Pop().value, n.value)
	testsupport.AssertThatInt(t, 0, s.Len())
}
