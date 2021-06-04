package niu2

import "testing"

/*
描述
牛牛特别喜欢数字7，他想知道，一个数减去7后是否刚好是2的幂次方，不过他不知道该怎么做，所以他想请你帮忙。
给定一个数字n，如果该数减去7后是2的幂次方，返回"YES"，否则，返回"NO"。
示例1
输入：
9
复制
返回值：
"YES"
复制
说明：
9-7=2，是2的幂次方。
备注：
8 \leq n \leq 10^98≤n≤10
9
*/
func TestSolve(t *testing.T) {
	ex := 9
	ret := solve(ex)
	t.Log(ret)
}
func solve(n int) string {
	i := n - 7
	for i > 2 {
		if i%2 == 0 {
			i = i / 2
		} else {
			return "NO"
		}
	}
	return "YES"
}
