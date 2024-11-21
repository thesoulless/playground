package graph

import (
	"github.com/thesoulles/playground/go/algo/stack"
)

type Edge struct {
	To     int
	Weight int
}

func DFS(graph [][]Edge, src int, needle int) ([]int, bool) {
	seen := make([]bool, len(graph))

	path := stack.New[int]()

	walk(graph, src, needle, seen, path)

	if path.Lenght == 0 {
		return nil, false
	}

	length := path.Lenght
	result := make([]int, length)
	for i := 0; i < length; i++ {
		v, err := path.Pop()
		if err != nil {
			panic(err)
		}

		result[length-i-1] = v
	}

	return result, true
}

func walk(graph [][]Edge, curr int, needle int, seen []bool, path *stack.Stack[int]) bool {
	if seen[curr] {
		return false
	}

	seen[curr] = true

	path.Push(curr)

	if curr == needle {
		return true
	}

	list := graph[curr]
	for _, edge := range list {
		if walk(graph, edge.To, needle, seen, path) {
			return true
		}
	}

	path.Pop()

	return false
}
