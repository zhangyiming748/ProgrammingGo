package neighbor

import (
	"math"
	"sort"
	"testing"
)

/*
描述
牛牛最近搬到了一座新的城镇，这个城镇可以看成是一个一维的坐标系。城镇上有n个居民，第i个居民的位置为a_ia
i
​
 。现在牛牛有m个搬家方案，在第i个方案中他会搬到位置x_ix
i
​
 。

俗话说的好，远亲不如近邻。现在牛牛想知道，对于每个搬家方案，搬家后与最近的居民的距离为多少。

示例1
输入：
3,2,[2,4,7],[5,8]
复制
返回值：
[1,1]
复制
说明：
第一个方案搬到位置5，与5最近的居民在位置4，距离为1.
第二个方案搬到位置8，与8最近的居民在位置7，距离为1
备注：
第一个参数为int型变量，代表居民个数n
第二个参数为int型变量，代表方案个数m
第三个参数为vector<int>，包含n个元素代表n个居民的位置
第四个参数为vector<int>，包含m个元素代表m个方案对应的位置
*/
func TestSolve(t *testing.T) {
	ret := solve(3, 2, []int{2, 4, 7}, []int{5, 8})
	t.Log(ret)
}
func solve(n int, m int, a []int, x []int) []int {
	sort.Ints(a)
	res := make([]int, 0)
	for i := 0; i < m; i++ {
		res[i] = getRes(a, x[i])
	}
	return res
}
func getRes(s []int, target int) int {
	c1 := math.Abs(float64(s[0] - target))
	c2 := math.Abs(float64(s[len(s)-1] - target))
	nearest := math.Min(c1, c2)
	l := 0
	r := s[len(s)-1]
	mid := (l + r) / 2
	min := math.Min(math.Abs(float64(s[mid]-target)), nearest)
	nearest = math.Abs(min)
	for l < r {
		if s[mid] > target {
			r = mid - 1
		} else if s[mid] < target {
			l = mid + 1
		} else {
			return 0
		}
		mid = (l + r) / 2
		min = math.Min(math.Abs(float64(s[mid]-target)), nearest)
		nearest = math.Abs(min)
	}
	abs := math.Abs(float64(s[l] - target))
	min = math.Min(abs, nearest)
	nearest = math.Abs(min)
	return int(nearest)
}
