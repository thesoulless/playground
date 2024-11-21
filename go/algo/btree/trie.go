package btree

type AutoTrieNode struct {
	end   bool
	char  rune
	child map[rune]*AutoTrieNode
}

func NewTrie() *AutoTrieNode {
	trie := &AutoTrieNode{}

	trie.child = map[rune]*AutoTrieNode{}

	return trie
}

func Insert(head *AutoTrieNode, str string) {
	curr := head
	for _, r := range str {
		if curr.child[r] != nil {
			curr = curr.child[r]
		} else {
			node := &AutoTrieNode{char: r}
			if curr.child == nil {
				curr.child = map[rune]*AutoTrieNode{}
			}
			curr.child[r] = node
			curr = node
		}
	}

	curr.end = true
}

func FindHeap(head *AutoTrieNode, str string) []string {
	curr := head
	for _, r := range str {
		if head.child[r] != nil {
			curr = head.child[r]
		}
	}

	vals := HeapDFS(curr)
	for i, v := range vals {
		vals[i] = str + v
	}

	return vals
}

func HeapDFS(head *AutoTrieNode) []string {
	if head == nil {
		return nil
	}

	var res []string
	if head.end {
		res = append(res, string(head.char))
	}

	for _, c := range head.child {
		if c != nil {
			suffixes := HeapDFS(c)

			for _, suffix := range suffixes {
				if c.end {
					res = append(res, suffix)
				} else {
					res = append(res, string(c.char)+suffix)
				}
			}
		}
	}

	return res
}
