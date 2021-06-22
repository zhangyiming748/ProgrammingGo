package convertHex

import (
	"strings"
	"testing"
)

/*
描述
给定一个十进制数M，以及需要转换的进制数N。将十进制数M转化为N进制数
示例1
输入：
7,2
复制
返回值：
"111"

*/

func solve(M int, N int) string {
	var res string
	var flag bool
	if M < 0 {
		M = -M
		flag = true
	}
	const a = "0123456789ABCDEF"
	for M != 0 {
		res = strings.Join([]string{string(a[M%N]), res}, "")
		M /= N
	}
	result := string(res)
	if flag {
		result = strings.Join([]string{"-", result}, "")
	}
	return result
}
func TestSolve(t *testing.T) {
	ret := solve(-4, 3)
	t.Log(ret)
}
