package search

import (
	"math"
)

func TwoCrystalBalls(breaks []bool) int {
	jumpAmount := int(math.Sqrt(float64(len(breaks))))

	i := jumpAmount
	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	j := 0
	for ; j < jumpAmount && i < len(breaks); j, i = j+1, i+1 {
		if breaks[i] {
			return i
		}
	}
	return -1
}
