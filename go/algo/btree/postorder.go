package btree

import "cmp"

func VisitPostOrder[T cmp.Ordered](head *Node[T]) []T {
	return visitPostOrder(head)
}

func visitPostOrder[T cmp.Ordered](curr *Node[T]) []T {
	if curr == nil {
		return nil
	}

	ns := visitPostOrder(curr.Left)

	ns = append(ns, visitPostOrder(curr.Right)...)
	ns = append(ns, curr.Value)

	return ns
}
