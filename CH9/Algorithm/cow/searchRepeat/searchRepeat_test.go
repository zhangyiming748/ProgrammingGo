package searchRepeat

import "testing"

/*
描述
在有 n+1 个数的序列 a 中找出重复的数

示例1
输入：
4,[1,2,1,4,3]
复制
返回值：
1
复制
备注：
其中1<=n<=100000。
要求时间复杂度为O(n)，额外空间复杂度为O(1)
*/
func TestSearch(t *testing.T) {
	ret := search(4, []int{4, 1, 2, 1, 4, 3})
	t.Log(ret)
}
func search(n int, a []int) int {
	m := make(map[int]bool)
	for _, v := range a {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v] = true
		}
	}
	return 0
}
