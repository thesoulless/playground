package search_test

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/thesoulles/playground/go/algo/search"
)

func TestTwoCrystalBalls(t *testing.T) {
	bindx, err := rand.Int(rand.Reader, big.NewInt(100_000))
	if err != nil {
		t.Errorf("%v", err)
	}

	index := bindx.Int64()

	data := make([]bool, 100_000)
	for i := range data {
		if int64(i) < index {
			data[i] = false
		} else {
			data[i] = true
		}
	}

	idx, found := search.TwoCrystalBall(data)
	if !found || int64(idx) != index {
		t.Errorf("result does  not match, got: (%d, %t), expected: (%d, %t)", idx, found, index, true)
	}

	data = make([]bool, 100_000)

	idx, found = search.TwoCrystalBall(data)
	if found {
		t.Errorf("result does  not match, got: (%d, %t), expected: (%d, %t)", idx, found, index, true)
	}
}
