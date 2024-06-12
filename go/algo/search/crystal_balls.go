package search

import (
	"math"
)

func TwoCrystalBall(S []bool) (int, bool) {
	n := len(S)
	l := 0

	if n == 0 {
		return -1, false
	}

	if S[l] {
		return 0, true
	}

	step := int(math.Sqrt(float64(n)))
	for ; l < n && !S[l]; l += step {
	}

	l -= step
	for j := 0; j < step && l < n && !S[l]; j++ {
		l++
	}

	return l, l < n && S[l]
}
