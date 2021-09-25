package datastructures

type Stack struct {
	values []*Node
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(node *Node) {
	s.values = append(s.values, node)
}

func (s *Stack) Pop() *Node {
	n := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return n
}

func (s *Stack) Len() int {
	return len(s.values)
}

func (s *Stack) isEmpty() bool {
	return len(s.values) == 0
}
