package graph

import (
	"slices"

	"github.com/thesoulles/playground/go/algo/btree"
)

func Dijkstra(graph [][]Edge, src int, sink int) []int {
	seen := make([]bool, len(graph))
	dists := make([]int, len(graph))
	prev := make([]int, len(graph))

	for i := range prev {
		prev[i] = -1
	}

	for i := range dists {
		dists[i] = 1 << 31
	}

	dists[src] = 0

	for hasUnseen(seen, dists) {
		curr := getMinUnseen(seen, dists)
		seen[curr] = true

		edges := graph[curr]
		for _, edge := range edges {
			if seen[edge.To] {
				continue
			}

			dist := dists[curr] + edge.Weight
			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	var res []int
	curr := sink

	for prev[curr] != -1 {
		res = append(res, curr)
		curr = prev[curr]
	}
	res = append(res, src)
	slices.Reverse(res)

	return res
}

func Dijkstra2(graph [][]Edge, src int, sink int) []int {
	dists := make([]int, len(graph))
	prev := make([]int, len(graph))

	mh := btree.NewMinHeap[int](len(graph))
	for i := range dists {
		dists[i] = 1 << 31
	}
	dists[src] = 0

	for v := range dists {
		mh.Insert(v)
	}

	for i := range prev {
		prev[i] = -1
	}

	for curr, _, err := mh.Delete(0); err == nil; curr, _, err = mh.Delete(0) {
		edges := graph[curr]
		for _, edge := range edges {
			dist := dists[curr] + edge.Weight
			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	var res []int
	curr := sink

	for prev[curr] != -1 {
		res = append(res, curr)
		curr = prev[curr]
	}
	res = append(res, src)
	slices.Reverse(res)

	return res
}

func hasUnseen(seen []bool, dists []int) bool {
	for i, v := range seen {
		if !v && dists[i] < (1<<31) {
			return true
		}
	}

	return false
}

func getMinUnseen(seen []bool, dists []int) int {
	idx := -1
	m := 1 << 31

	for i, s := range seen {
		if !s && dists[i] < m {
			m = dists[i]
			idx = i
		}
	}

	return idx
}
