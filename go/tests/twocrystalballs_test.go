// tests/twocrystalballs_test.go
package search_test

import (
	"algorithm_katas/algorithms/search"
	"math/rand"
	"testing"
	"time"
)

func TestTwoCrystalBalls(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	idx := rand.Intn(10000)
	data := make([]bool, 10000)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	if search.TwoCrystalBalls(data) != idx {
		t.Errorf("Expected %d, got %d", idx, search.TwoCrystalBalls(data))
	}

	data = make([]bool, 821)
	if search.TwoCrystalBalls(data) != -1 {
		t.Errorf("Expected -1, got %d", search.TwoCrystalBalls(data))
	}
}
