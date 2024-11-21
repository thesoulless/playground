package graph

import (
	"testing"
)

func TestBFS(t *testing.T) {
	graph := WeightedAdjMatrix{
		// 0  1  2  3
		/*{0, 1, 1, 0}, // 0
		{1, 0, 0, 1}, // 1
		{1, 0, 0, 1}, // 2
		{0, 1, 1, 0}, // 3*/
		{0, 3, 1, 0, 0, 0, 0}, // 0
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 5, 0, 2, 0},
		{0, 0, 18, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 1},
	}

	src := 0
	needle := 6

	path, found := BFS(graph, src, needle)
	if !found {
		t.Errorf("Expected path to be found, got %v", found)
	}

	expected := []int{0, 1, 4, 5, 6}
	for i, v := range path {
		if v != expected[i] {
			t.Errorf("Expected path to be %v, got %v", expected, path)
			break
		}
	}
}
