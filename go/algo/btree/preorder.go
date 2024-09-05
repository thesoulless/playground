package bree

import "cmp"

type Node[T cmp.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func VisitPreOrder[T cmp.Ordered](head *Node[T]) []T {
	return visitPreOrder(head)
}

func visitPreOrder[T cmp.Ordered](curr *Node[T]) []T {
	if curr == nil {
		return nil
	}

	ns := []T{curr.Value}

	ns = append(ns, visitPreOrder(curr.Left)...)
	ns = append(ns, visitPreOrder(curr.Right)...)

	return ns
}
