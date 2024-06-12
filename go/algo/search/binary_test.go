package search_test

import (
	"testing"

	"github.com/thesoulles/playground/go/algo/search"
)

func TestBinary(t *testing.T) {
	type iArgs struct {
		haystack []int
		needle   int
	}

	iTests := []struct {
		name  string
		args  iArgs
		found bool
		index int
	}{
		{
			name:  "empty list",
			args:  iArgs{},
			found: false,
			index: -1,
		},
		{
			name: "first",
			args: iArgs{
				haystack: []int{2, 3, 6, 9},
				needle:   2,
			},
			found: true,
			index: 0,
		},
		{
			name: "last",
			args: iArgs{
				haystack: []int{2, 3, 6, 9},
				needle:   9,
			},
			found: true,
			index: 3,
		},
		{
			name: "mid",
			args: iArgs{
				haystack: []int{2, 3, 6, 9},
				needle:   6,
			},
			found: true,
			index: 2,
		},
		{
			name: "not found",
			args: iArgs{
				haystack: []int{2, 3, 6, 9},
				needle:   12,
			},
			found: false,
			index: -1,
		},
	}

	for _, tt := range iTests {
		t.Run(tt.name, func(t *testing.T) {
			i, f := search.Binary(tt.args.haystack, tt.args.needle)

			if tt.found && i != tt.index {
				t.Errorf("index does not match. Got: %d, expected: %d", i, tt.index)
			}

			if f != tt.found {
				t.Errorf("%s: found does not match. Got: %t, expected: %t", tt.name, f, tt.found)
			}
		})
	}
}
