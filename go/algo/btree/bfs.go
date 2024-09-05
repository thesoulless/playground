package bree

import (
	"cmp"

	"github.com/thesoulles/playground/go/algo/queue"
)

func BFS[T cmp.Ordered](head *Node[T], needle T) bool {
	if head == nil {
		return false
	}

	return bfs(head, needle)
}

func bfs[T cmp.Ordered](head *Node[T], needle T) bool {
	q := queue.New[*Node[T]]()
	q.Enqueue(head)

	for q.Length() > 0 {
		curr, err := q.Dequeue()
		if err != nil {
			return false
		}

		if curr.Value == needle {
			return true
		}

		if curr.Left != nil {
			q.Enqueue(curr.Left)
		}

		if curr.Right != nil {
			q.Enqueue(curr.Right)
		}
	}

	return false
}
