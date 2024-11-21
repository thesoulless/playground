package btree_test

import (
	"testing"

	btree "github.com/thesoulles/playground/go/algo/btree"
)

func TestVisitInOrder(t *testing.T) {
	head := &btree.Node[int]{}
	head.Value = 1
	head.Left = &btree.Node[int]{}
	head.Left.Value = 2
	head.Right = &btree.Node[int]{}
	head.Right.Value = 3

	got := btree.VisitInorder(head)
	want := []int{2, 1, 3}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
