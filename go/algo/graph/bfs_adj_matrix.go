package graph

import (
	"github.com/thesoulles/playground/go/algo/queue"
	"github.com/thesoulles/playground/go/algo/stack"
)

type WeightedAdjMatrix [][]int

func BFS(graph WeightedAdjMatrix, src int, needle int) ([]int, bool) {
	prev := make([]int, len(graph))
	seen := make([]bool, len(graph))

	for i := range prev {
		prev[i] = -1
	}

	seen[0] = true

	q := queue.New[int]()
	q.Enqueue(src)

	for q.Length() > 0 {
		curr, err := q.Dequeue()
		if err != nil {
			panic(err)
		}

		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i, v := range adjs {
			if v == 0 {
				continue
			}

			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q.Enqueue(i)
		}
		seen[curr] = true
	}

	if prev[needle] == -1 {
		return nil, false
	}

	curr := needle
	resq := stack.New[int]()

	res := []int{src}

	for prev[curr] != -1 {
		resq.Push(curr)
		// res = append(res, curr)
		curr = prev[curr]
	}

	for resq.Length() > 0 {
		curr, err := resq.Pop()
		if err != nil {
			panic(err)
		}
		res = append(res, curr)
	}

	return res, true
}
