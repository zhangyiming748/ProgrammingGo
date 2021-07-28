package reverse

import (
	"math"
	"strconv"
	"testing"
)

func reverse(x int) int {
	nagi := nagitive(x)
	if nagi {
		x = -x
	}
	xs := strconv.Itoa(x)
	xb := []byte(xs)
	nb := []byte{}
	for i := len(xb) - 1; i >= 0; i-- {
		nb = append(nb, xs[i])
	}
	sb := string(nb)
	ret, _ := strconv.Atoi(sb)
	if nagi {
		ret = -ret
	}
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	}
	return ret
}
func nagitive(i int) bool {
	if i < 0 {
		return true
	}
	return false
}
func TestReverse(t *testing.T) {
	x, y := 123, -321
	ret1 := reverse(x)
	ret2 := reverse(y)
	t.Logf("ret1=%v\tret2=%v\n", ret1, ret2)
}
