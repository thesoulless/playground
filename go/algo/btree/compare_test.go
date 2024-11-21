package btree_test

import (
	"testing"

	btree "github.com/thesoulles/playground/go/algo/btree"
)

func TestCompare(t *testing.T) {
	// Compare two trees
	//     1
	//    / \
	//   2   3
	//  / \
	// 4   5
	// and
	//     1
	//    / \
	//   2   3
	//  / \
	// 4   5
	// they are the same
	head1 := &btree.Node[int]{Value: 1}
	head1.Left = &btree.Node[int]{Value: 2}
	head1.Right = &btree.Node[int]{Value: 3}
	head1.Left.Left = &btree.Node[int]{Value: 4}
	head1.Left.Right = &btree.Node[int]{Value: 5}

	head2 := &btree.Node[int]{Value: 1}
	head2.Left = &btree.Node[int]{Value: 2}
	head2.Right = &btree.Node[int]{Value: 3}
	head2.Left.Left = &btree.Node[int]{Value: 4}
	head2.Left.Right = &btree.Node[int]{Value: 5}

	if !btree.Compare(head1, head2) {
		t.Error("Expected true, got false")
	}

	//       1
	//      / \
	//     2   3
	//    /
	//   4
	//  /
	// 5
	// and
	//     1
	//    / \
	//   2   3
	//  / \
	// 4   5
	head1 = &btree.Node[int]{Value: 1}
	head1.Left = &btree.Node[int]{Value: 2}
	head1.Right = &btree.Node[int]{Value: 3}
	head1.Left.Left = &btree.Node[int]{Value: 4}
	head1.Left.Left.Left = &btree.Node[int]{Value: 5}

	head2 = &btree.Node[int]{Value: 1}
	head2.Left = &btree.Node[int]{Value: 2}
	head2.Right = &btree.Node[int]{Value: 3}
	head2.Left.Left = &btree.Node[int]{Value: 4}
	head2.Left.Right = &btree.Node[int]{Value: 5}

	if btree.Compare(head1, head2) {
		t.Error("Expected false, got true")
	}
}
