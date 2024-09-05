package bree_test

import (
	"testing"

	btree "github.com/thesoulles/playground/go/algo/btree"
)

func TestVisitPreOrder(t *testing.T) {
	head := &btree.Node[int]{}
	head.Value = 1
	head.Left = &btree.Node[int]{}
	head.Left.Value = 2
	head.Right = &btree.Node[int]{}
	head.Right.Value = 3

	got := btree.VisitPreOrder(head)
	want := []int{1, 2, 3}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
