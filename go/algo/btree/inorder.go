package btree

import "cmp"

// DFS
// DFS preserves the shape of the tree
// @TODO: use a stack

func VisitInorder[T cmp.Ordered](head *Node[T]) []T {
	return visitInorder(head)
}

func visitInorder[T cmp.Ordered](curr *Node[T]) []T {
	if curr == nil {
		return nil
	}

	ns := visitInorder(curr.Left)

	ns = append(ns, curr.Value)
	ns = append(ns, visitInorder(curr.Right)...)

	return ns
}
