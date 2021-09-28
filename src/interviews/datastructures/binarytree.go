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
// output = 1,2,3,4,5
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

// Left -> Right -> Root
// output = 1,3,2,5,4
func PostorderTraversal(node *Node) {
	s1 := newStack()
	s2 := newStack()

	s1.Push(node)

	var temp *Node

	for {
		if s1.Len() == 0 {
			break
		}
		temp = s1.Pop()
		s2.Push(temp)

		if temp.left != nil {
			s1.Push(temp.left)
		}
		if temp.right != nil {
			s1.Push(temp.right)
		}
	}

	for {
		if s2.isEmpty() {
			break
		}
		log.Print(s2.Pop().value)
	}

}

// root -> left -> right
func PreorderTraversal(node *Node) {
	for {
		if node == nil {
			break
		}

		if node.left == nil {
			log.Print(node.value)
			node = node.right
		} else {
			c := node.left

			for {
				if c.right != nil && c.right != node {
					c = c.right
				} else {
					break
				}
			}

			if c.right == node {
				c.right = nil
				node = node.right
			} else {
				log.Print(node.value)
				c.right = node
				node = node.left
			}
		}

	}
}

// ** NOTES **
// All recursion is simpy using stack and values so can be converted to iteration and is simpler for interviews -
// lambda calculus (rec) vs turing machine (iter) -> church-turing thesis
