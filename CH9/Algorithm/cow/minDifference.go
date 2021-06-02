package cow

import "math"

func minDifference(a []int) int {
	min := math.MaxInt32
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			sub := a[i] - a[j]
			if sub < 0 {
				sub = 0 - sub
			}
			if sub < min {
				min = sub
			}
		}
	}
	return min
}
