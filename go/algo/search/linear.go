package search

import (
	"cmp"
)

func Linear[S ~[]E, E cmp.Ordered](haystack S, needle E) (int, bool) {
	if len(haystack) == 0 {
		return -1, false
	}

	for i, e := range haystack {
		if e == needle {
			return i, true
		}
	}

	return -1, false
}
