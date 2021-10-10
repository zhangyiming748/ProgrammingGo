package NumberOf1

import (
	"fmt"
	"testing"
)

/*
输入一个整数，输出该数32位二进制表示中1的个数。其中负数用补码表示。
*/
func NumberOf1(n int) int {
	b := fmt.Sprintf("%b", n)
	var nagitive bool
	if n<0{
		nagitive = true
		n = -n
	}
	var count int
	fmt.Printf("%v %v",b,nagitive)
	for _, v := range b {
		if v == '1' {
			count++
		}
	}
	return count
}
func TestUnit(t *testing.T) {
	ret := NumberOf1(-1)
	t.Log(ret)
}
