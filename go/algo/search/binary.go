package search

import (
	"cmp"
)

// Binary search on an ordered haystack
// **Note:** this does not handle isNaN shenanigans for floats
func Binary[S ~[]E, E cmp.Ordered](haystack S, needle E) (int, bool) {
	n := len(haystack)

	if n == 0 {
		return -1, false
	}

	l, h := 0, n

	for l < h {
		m := int(uint(l+h) >> 1)

		if cmp.Less(haystack[m], needle) {
			l = m + 1
		} else {
			h = m
		}
	}

	return l, l < n && haystack[l] == needle
}
