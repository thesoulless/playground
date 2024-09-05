package bree_test

import (
	"testing"

	btree "github.com/thesoulles/playground/go/algo/btree"
)

func Test_BFS(t *testing.T) {
	head1 := &btree.Node[int]{}
	head1.Value = 1
	head1.Left = &btree.Node[int]{}
	head1.Left.Value = 2
	head1.Right = &btree.Node[int]{}
	head1.Right.Value = 3

	tests := []struct {
		name   string
		tree   *btree.Node[int]
		needle int
		want   bool
	}{
		{
			name:   "should find",
			tree:   head1,
			needle: 2,
			want:   true,
		},
		{
			name:   "nil",
			tree:   nil,
			needle: 2,
			want:   false,
		},
		{
			name:   "no values",
			tree:   &btree.Node[int]{},
			needle: 2,
			want:   false,
		},
		{
			name:   "head",
			tree:   head1,
			needle: 1,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := btree.BFS(tt.tree, tt.needle)
			if got != tt.want {
				t.Errorf("with %v\nexpected: %t, got: %t", tt.tree, got, tt.want)
			}
		})
	}
}
