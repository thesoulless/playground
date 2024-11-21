package btree

import "cmp"

// bst is a binary search tree.

func Find[T cmp.Ordered](head *Node[T], needle T) bool {
	if head == nil {
		return false
	}

	if needle == head.Value {
		return true
	}

	if needle <= head.Value {
		return Find(head.Left, needle)
	}

	return Find(head.Right, needle)
}

func InsertBST[T cmp.Ordered](head *Node[T], needle T) *Node[T] {
	if head == nil {
		return &Node[T]{Value: needle}
	}

	if needle <= head.Value {
		return InsertBST(head.Left, needle)
	} else {
		return InsertBST(head.Right, needle)
	}
}

func deleteBST[T cmp.Ordered](head *Node[T], needle T, parent *Node[T]) bool {
	if head == nil {
		return false
	}

	if needle == head.Value {
		if head.Right == nil && head.Left == nil {
			newHead := head.Left
			for {
				nh := newHead.Right
				if nh.Right == nil {
					removeBSTNode(newHead, nh)
					newHead = nh
					break
				}
				newHead = nh
			}

			if parent != nil {
				if needle <= parent.Value {
					parent.Left = newHead
				} else {
					parent.Right = newHead
				}
			}

			return true
		} else {
			if head.Right != nil {
				removeBSTNode(head, head.Right)
			} else {
				removeBSTNode(head, head.Left)
			}

			return true
		}
	}

	if needle <= head.Value {
		return deleteBST(head.Left, needle, head)
	} else {
		return deleteBST(head.Right, needle, head)
	}
}

func removeBSTNode[T cmp.Ordered](head *Node[T], node *Node[T]) {
	sub := node.Right
	if node.Right == nil {
		sub = node.Left
	}

	if head.Right != nil {
		head.Right = sub
	} else {
		head.Left = sub
	}
}

func DeleteBST[T cmp.Ordered](head *Node[T], needle T) bool {
	return deleteBST(head, needle, nil)
}
