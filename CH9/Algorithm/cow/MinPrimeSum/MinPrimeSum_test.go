package MinPrimeSum

import (
	"testing"
)

/*
描述
现在给定一个正整数N，牛牛希望知道N最少表示成多少个素数的和。
素数是指在大于1的自然数中，除了1和它本身以外不再有其他因数的自然数。

提示
哥德巴赫猜想：任意大于2的偶数都可以拆分成两个质数之和。该猜想尚未严格证明，但暂时没有找到反例。

示例1
输入：
3
复制
返回值：
1
复制
说明：
3本身就是1个素数
示例2
输入：
6
复制
返回值：
2
复制
说明：
6可以表示为3 + 3，注意同样的素数可以使用多次
*/
func TestMinPrimeSum(t *testing.T) {
	ex1 := 3
	ex2 := 6
	ret1 := MinPrimeSum(ex1)
	ret2 := MinPrimeSum(ex2)
	t.Logf("ret1=%d\tret2=%d", ret1, ret2)
}

func MinPrimeSum(N int) int {
	res := 0
	if N == 0 || N == 1 {
		return res
	}
	if N%2 == 0 {
		if N == 2 {
			res = 1
		} else {
			res = 2
		}
	} else {
		if isSU(N) {
			res = 1
		} else {
			if isSU(N - 2) {
				res = 2
			} else {
				res = 3
			}

		}
	}
	return res
}
func isSU(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
