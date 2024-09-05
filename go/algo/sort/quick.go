package sort

import "cmp"

func partition[S ~[]E, E cmp.Ordered](s S, lo, hi int) int {
	pivot := s[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if s[j] <= pivot {
			i++
			s[i], s[j] = s[j], s[i]
		}
	}

	i++
	s[i], s[hi] = s[hi], s[i]

	return i
}

func qs[S ~[]E, E cmp.Ordered](s S, lo, hi int) {
	if lo < hi {
		p := partition(s, lo, hi)
		qs(s, lo, p-1)
		qs(s, p+1, hi)
	}
}

func Quick[S ~[]E, E cmp.Ordered](s S) {
	qs(s, 0, len(s)-1)
}
