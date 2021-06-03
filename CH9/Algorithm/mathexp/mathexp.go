package mathexp
/*
描述
给出一个数字n，需要不断地将所有数位上的值做乘法运算，直至最后数字不发生变化为止。
问最后生成的数字为多少？
示例1
输入：
10
复制
返回值：
0
复制
示例2
输入：
55
复制
返回值：
0
复制
说明：
55 -> 5 * 5 = 25 -> 2 * 5 = 10 -> 1 * 0 = 0
备注：
1\leq n\leq 10^{18}1≤n≤10
18
*/
func mathexp(n int64)int {
	var  single int64 =0
	var res int64=1
	if n<10{
		return int(n)
	}
	for n>=1{
		single=n%10
		n=n/10
		res = single*res
	}
	return mathexp(res)
}
