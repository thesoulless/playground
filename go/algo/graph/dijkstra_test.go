package graph

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := make([][]Edge, 7)
	graph[0] = []Edge{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	}
	graph[1] = []Edge{
		{To: 0, Weight: 3},
		{To: 2, Weight: 4},
		{To: 4, Weight: 1},
	}
	graph[2] = []Edge{
		{To: 1, Weight: 4},
		{To: 3, Weight: 7},
		{To: 0, Weight: 1},
	}
	graph[3] = []Edge{
		{To: 2, Weight: 7},
		{To: 4, Weight: 5},
		{To: 6, Weight: 1},
	}
	graph[4] = []Edge{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	}
	graph[5] = []Edge{
		{To: 6, Weight: 1},
		{To: 4, Weight: 2},
		{To: 2, Weight: 18},
	}
	graph[6] = []Edge{
		{To: 3, Weight: 1},
		{To: 5, Weight: 1},
	}

	res := Dijkstra2(graph, 0, 6)

	expected := []int{0, 1, 4, 5, 6}

	if len(res) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, res)
	}

	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, res)
		}
	}
}
