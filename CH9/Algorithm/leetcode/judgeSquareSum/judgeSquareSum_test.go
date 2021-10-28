package judgeSquareSum

import (
	"math"
	"testing"
)

//给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。
func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}
	return false
}

func TestUnit(t *testing.T) {
	var (
		a int = 5
		b int = 3
		c int = 4
		d int = 2
		e int = 1
	)
	ret1 := judgeSquareSum(a)
	ret2 := judgeSquareSum(b)
	ret3 := judgeSquareSum(c)
	ret4 := judgeSquareSum(d)
	ret5 := judgeSquareSum(e)
	t.Log(ret1, ret2, ret3, ret4, ret5)
}
