package splitArray

import (
	"sort"
	"testing"
)

/*
描述
牛牛面前有一堆数，他想把这些数分成两堆，只不过牛牛是一个很有想法的人。
他希望分得的两堆数能够满足，第一堆数的最大值和第二堆数的最小值差值最小。
由于数太多，牛牛犯了难，所以他想请你帮帮他，给定n个数，返回符合牛牛希望的分法中最小的差值是多少。
示例1
输入：
2,[1,2] 返回值：1
只有一种分发，分成的两堆数中第一堆数的最大值与第二堆数的最小值差值最小为1。
示例2
输入：
6,[2,4,3,3,1,2]
复制
返回值：
0
复制
说明：
在所有分法中，第一堆数分成[2,1]，第二堆数分成[3,2,4,3]，分成的两堆数中第一堆数的最大值与第二堆数的最小值差值最小为0。
*/
func TestSolution(t *testing.T) {
	n := 6
	array := []int{2, 4, 3, 3, 1, 2}
	res := splitArray(n, array)
	t.Log(res)
}
func splitArray(n int, a []int) int {
	sort.Ints(a)
	res := a[1] - a[0]
	for i := 0; i < n-1; i++ {
		res = min(res, a[i+1]-a[i])
	}
	return res
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
