package btree

import "testing"

func TestHeapDFS(t *testing.T) {
	//         []
	//      [c]
	//   [a]
	// [t] [r]
	head := NewTrie()
	Insert(head, "cat")
	Insert(head, "car")

	curr := head
	curr = curr.child['c']

	// HeapDFS(head *AutoTrieNode)
	// dumpTree(t, head)

	res := HeapDFS(curr)
	t.Logf("%#v", res)
}

func dumpTree(t *testing.T, head *AutoTrieNode) {
	t.Helper()
	// t.Logf("test\n")

	if head.char != 0 {
		t.Logf("%c", head.char)
	}

	for _, c := range head.child {
		if c != nil {
			dumpTree(t, c)
		}
	}

}
