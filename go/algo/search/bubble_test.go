package search_test

import (
	"crypto/rand"
	"math"
	"math/big"
	"slices"
	"testing"

	"github.com/thesoulles/playground/go/algo/search"
)

func TestBubble(t *testing.T) {
	r0 := randIntArray(t, 0)
	r1 := randIntArray(t, 1)
	r2 := randIntArray(t, 5)
	r3 := randIntArray(t, 100)
	r4 := randIntArray(t, 1000)

	iTests := []struct {
		name string
		s    []int
	}{
		{
			name: "zero elements",
			s:    r0,
		},
		{
			name: "one element",
			s:    r1,
		},
		{
			name: "five elements",
			s:    r2,
		},
		{
			name: "100 elements",
			s:    r3,
		},
		{
			name: "1000 elements",
			s:    r4,
		},
	}

	for _, tt := range iTests {
		t.Run(tt.name, func(t *testing.T) {
			sc := make([]int, len(tt.s))
			_ = copy(sc, tt.s)
			search.Bubble(tt.s)

			slices.Sort(sc)

			for i := range tt.s {
				if tt.s[i] != sc[i] {
					t.Errorf("sorted arrays are not equal. got: \n%v\nExpected: \n%v\n", tt.s, sc)
					t.FailNow()
				}
			}
		})
	}
}

func randIntArray(t *testing.T, n int) []int {
	t.Helper()

	if n <= 0 {
		n = 100
	}

	s := make([]int, n)

	for i := range s {
		br, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
		if err != nil {
			t.Errorf("failed to generate random number: %v", err)
		}

		s[i] = int(br.Int64())
	}

	return s
}
