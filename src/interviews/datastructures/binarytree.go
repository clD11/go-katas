package datastructures

import "log"

// Binary Trees - node can have at most two children

// Binary Search Tree - the key in each node is greater than or equal to any key stored in the left sub-tree, and less than or equal to any key stored in the right sub-tree
//			4
//		2		5
//	1		3

// Operations

func Insert(node *Node) {

}

func Delete(node *Node) {

}

func Search(node *Node) {

}

// Traversal (stack)

// Inorder (Left, Root, Right): 1,2,3,4,5
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

// Preorder (Root, Left, Right): 4,2,1,3,5
func PreorderTraversal(node *Node) {
	s := newStack()
	s.Push(node)

	for {
		if s.isEmpty() {
			break
		}
		current := s.Pop()
		log.Print(current.value)

		if current.right != nil {
			s.Push(current.right)
		}
		if current.left != nil {
			s.Push(current.left)
		}
	}
}

// Postorder (Left, Right, Root): 1,3,2,5,4
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

// ** NOTES **
// All recursion is simply using stack and values so can be converted to iteration and is simpler for interviews -
// lambda calculus (rec) vs turing machine (iter) -> church-turing thesis
