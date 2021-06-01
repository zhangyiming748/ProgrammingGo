package cow

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	var negative bool
	if x < 0 {
		negative = true
		x = 0 - x
	}
	xs := strconv.Itoa(x)
	sb := antiString(xs)

	ret, _ := strconv.Atoi(sb)
	if ret > math.MaxInt32 {
		ret = 0
	}
	if negative {
		ret = 0 - ret
	}
	return ret
}
