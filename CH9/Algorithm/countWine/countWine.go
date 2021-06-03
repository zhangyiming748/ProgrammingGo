package countWine

/*
描述
一瓶酒m元钱，两个酒瓶可以换一瓶酒，四个瓶盖可以换一瓶酒，现在有 n 元钱，求最多可以喝多少瓶酒？
（注：没有借贷功能，即最终不允许借一瓶酒、喝完后拿酒瓶兑换归还的操作）
示例1
输入：
2,12
复制
返回值：
19
复制
说明：
酒鬼总计可以喝19瓶酒
备注：
0 < m < 100
0 < n < 2000
*/
func countWine(m, n int) int {
	x := n / m
	var i int = x / 2
	j := x % 2
	var p int = x / 4
	q := x % 4
	sum := x + i + p
	for i > 0 || p > 0 {
		tmp1 := i
		tmp2 := p
		i = (i + j + tmp2) / 2
		j = (tmp1 + j + tmp2) % 2
		p = (p + q + tmp1) / 4
		q = (tmp2 + q + tmp1) % 4
		sum = sum + i + p
	}
	return sum
}
