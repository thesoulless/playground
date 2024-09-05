package bree

import "cmp"

// DFS

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
