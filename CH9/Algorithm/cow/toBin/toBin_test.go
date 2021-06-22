package toBin

import (
	"fmt"
	"testing"
)

/*
描述
牛牛想把一个数n转化为八位的二进制数，只不过牛牛不知道该怎么做，所以他想请你帮忙。
给定一个数n，返回将这个数转化为八位的二进制数(不足八位，往前补0)。
示例1
输入：
1
复制
返回值：
"00000001"
复制
备注：
0 \leq n \leq 2550≤n≤255
*/
func tranBinary(n int) string {
	b := fmt.Sprintf("%08b", n)
	return b
}
func TestTranBinary(t *testing.T) {
	in := 3
	ret := tranBinary(in)
	t.Log(ret)
}
