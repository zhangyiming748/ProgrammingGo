package multiBig

import (
	"math"
	"testing"
)

/*
给定一个无序数组，包含正数、负数和0，要求从中找出3个数的乘积，使得乘积最大，要求时间复杂度：O(n)，空间复杂度：O(1)。
输入：
[3,4,1,2]
返回值：
24
*/
func TestSolve(t *testing.T) {
	ret := solve([]int{3, 4, 1, 2})
	t.Log(ret)
}
func solve(a []int) int64 {
	var (
		max1, max2, max3 = math.MinInt32, math.MinInt32, math.MinInt32
		min1, min2       = math.MaxInt32, math.MaxInt32
	)
	for _, num := range a {
		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num > max2 {
			max3 = max2
			max2 = num
		} else if num > max3 {
			max3 = num
		}
		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}
	}
	biggest := max(int64(max1*max2*max3), int64(max1*min1*min2))
	return biggest
}
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
