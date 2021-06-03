package findNum

/*
描述
牛牛有两个数a和b，他想找到一个大于a且为b的倍数的最小整数，只不过他算数没学好，不知道该怎么做，现在他想请你帮忙。
给定两个数a和b，返回大于a且为b的倍数的最小整数。
示例1
输入：
3,2
复制
返回值：
4
复制
说明：
大于3且为2的倍数的最小整数为4。
备注：
1 \leq a,b \leq 10^91≤a,b≤10
9
*/
func findNumber(a int, b int) int {
	if a == 0 {
		return b
	}
	beat := 0
	for {
		if beat*b > a {
			return beat * b
		} else {
			beat++
		}
	}
}
