package datastructures

import (
	"testing"
)

func TestShouldRecurse(t *testing.T) {
	binarySearchTree := &Node{
		value: 4,
		left: &Node{
			value: 2,
			left: &Node{
				value: 1,
				left:  nil,
				right: nil,
			},
			right: &Node{
				value: 3,
				left:  nil,
				right: nil,
			},
		},
		right: &Node{
			value: 5,
			left:  nil,
			right: nil,
		},
	}
	InorderTraversal(binarySearchTree)
}
