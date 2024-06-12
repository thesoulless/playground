package search

import (
	"cmp"
)

func Bubble[S ~[]E, E cmp.Ordered](s S) {
	n := len(s)

	if n < 2 {
		return
	}

	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if cmp.Less(s[j+1], s[j]) {
				t := s[j]
				s[j] = s[j+1]
				s[j+1] = t
			}
		}
	}
}
