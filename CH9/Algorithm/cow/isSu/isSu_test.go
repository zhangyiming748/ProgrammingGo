package isSu

import "testing"

/*
描述
牛牛特别喜欢数字7，他想知道如果一个数字n乘以7是否是一个素数。
给定一个数字n，如果该数乘以7是一个素数，返回"YES"，否则，返回"NO"。
示例1
输入：
1
复制
返回值：
"YES"
复制
备注：
1 \leq n \leq 10^81≤n≤10
8
*/
func TestIsSu(t *testing.T) {
	num := 1
	ret := solve(num)
	t.Log(ret)
}
func solve(n int) string {
	if n == 1 {
		return "YES"
	}
	targe := n * 7
	for i := 2; i < targe; i++ {
		if targe%i == 0 {
			return "NO"
		}
	}
	return "YES"
}
