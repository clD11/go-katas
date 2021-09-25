package datastructures

import "log"

// Binary Trees

// Binary Search Tree
// The key in each Node is greater than or equal to any key stored in the left sub-tree, and less than or equal to any key stored in the right sub-tree
//			4
//		2		5
//	1		3

// Inorder
// Postorder
// Preorder

// left -> root -> right
func InorderTraversal(node *Node) {
	s := newStack()
	s.Push(node)

	current := node

	for {
		if s.Len() == 0 {
			break
		}

		// visit all left
		for {
			if current == nil {
				break
			}
			s.Push(current)
			current = current.left
		}

		// back track and go right (if right has a left then loop above goes down left and so on until all nodes are visited)
		current = s.Pop()

		// testing only isEmpty stop root being printed twice
		if !s.isEmpty() {
			log.Print(current.value)
		}

		current = current.right
	}
}

// ** NOTES **
// All recursion is simpy using stack and values so can be converted to iteration and is simpler for interviews -
// lambda calculus (rec) vs turing machine (iter) -> church-turing thesis
