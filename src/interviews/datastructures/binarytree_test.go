package datastructures

import (
	"testing"
)

func TestInOrderTraversal(t *testing.T) {
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

func TestPostOrderTraversal(t *testing.T) {
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
	PostorderTraversal(binarySearchTree)
}

func TestPreOrderTraversal(t *testing.T) {
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
	PreorderTraversal(binarySearchTree)
}
