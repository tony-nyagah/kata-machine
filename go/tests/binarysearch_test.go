// tests/binarysearch_test.go
package search_test

import (
	"algorithm_katas/algorithms/search"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	haystack := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	tests := []struct {
		name   string
		needle int
		want   bool
	}{
		{"find 69", 69, true},
		{"not find 1336", 1336, false},
		{"find 69420", 69420, true},
		{"not find 69421", 69421, false},
		{"find 1", 1, true},
		{"not find 0", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search.BinarySearch(haystack, tt.needle); got != tt.want {
				t.Errorf("BinarySearch(%v, %v) = %v, want %v", haystack, tt.needle, got, tt.want)
			}
		})
	}
}
