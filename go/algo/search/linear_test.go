package search_test

import (
	"testing"

	"github.com/thesoulles/playground/go/algo/search"
)

func TestLinear(t *testing.T) {
	type iArgs struct {
		haystack []int
		needle   int
	}

	intTests := []struct {
		name       string
		args       iArgs
		shouldFind bool
		onIndex    int
	}{
		{
			name:       "empty",
			shouldFind: false,
			onIndex:    -1,
		},
		{
			name: "first",
			args: iArgs{
				haystack: []int{23, 10, 4, 1, 232},
				needle:   23,
			},
			shouldFind: true,
			onIndex:    0,
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			i, ok := search.Linear(tt.args.haystack, tt.args.needle)

			if ok != tt.shouldFind {
				t.Errorf("find result does not match, got: %t, expexted: %t", ok, tt.shouldFind)
			}

			if i != tt.onIndex {
				t.Errorf("index does not match, got: %d, expected: %d", tt.onIndex, i)
			}
		})
	}
}
