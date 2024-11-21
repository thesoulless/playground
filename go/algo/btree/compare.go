package btree

import "cmp"

func Compare[T cmp.Ordered](head1 *Node[T], head2 *Node[T]) bool {

	if head1 == nil && head2 == nil {
		return true
	}

	if head1 == nil || head2 == nil {
		return false
	}

	if head1.Value != head2.Value {
		return false
	}

	return Compare(head1.Left, head2.Left) && Compare(head1.Right, head2.Right)
}
