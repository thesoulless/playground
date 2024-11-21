package graph

import (
	"testing"
)

func TestDFS(t *testing.T) {
	graph := make([][]Edge, 7)
	graph[0] = []Edge{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	}
	graph[1] = []Edge{
		{To: 4, Weight: 1},
	}
	graph[2] = []Edge{
		{To: 3, Weight: 7},
	}
	graph[3] = []Edge{}
	graph[4] = []Edge{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	}
	graph[5] = []Edge{
		{To: 2, Weight: 18},
		{To: 6, Weight: 1},
	}
	graph[6] = []Edge{
		{To: 3, Weight: 1},
	}

	src := 0
	needle := 6

	path, found := DFS(graph, src, needle)
	if !found {
		t.Errorf("Expected path to be found, got %v", found)
		t.Log(path)
	}

	expected := []int{0, 1, 4, 5, 6}
	for i, v := range path {
		if v != expected[i] {
			t.Errorf("Expected path to be %v, got %v", expected, path)
			break
		}
	}

	src = 6
	needle = 0

	path, found = DFS(graph, src, needle)
	if found {
		t.Errorf("Expected path to be found, got %v", found)
		t.Log(path)
	}

	if path != nil {
		t.Errorf("Expected path to be %v, got %v", expected, path)
	}
}
