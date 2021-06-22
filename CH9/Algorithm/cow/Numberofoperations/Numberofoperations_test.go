package Numberofoperations

import "testing"

/*
描述
给你一个正整数n，重复进行以下操作：
1.如果n是奇数，令n=n-3n=n−3
2.如果n是偶数，令n=n/2n=n/2
重复上述直至n=0停止，请输出进行操作的次数，如果n永远无法变成零，输出-1

示例1
输入：
2
复制
返回值：
-1
复制
说明：
1:2->1(2/2=1)
2:1->-2(1-3=-2)
3:-2->-1((-2)/2=(-1))
4.-1->-4(-1-3=-4)
5.-4->-2((-4)/2=(-2))
6.-2->-1((-2)/2=(-1))
.......开始进入第三步操作到第五步操作的循环，n永远无法等于0，所以返回-1。

示例2
输入：
9
复制
返回值：
3
复制
说明：
1.9->6(9-3=6)
2.6->3(6/2=3)
3.3->0(3-3=0)
三步操作后n变为0，所以返回3。

*/
func Numberofoperations(n int64) int {
	var i int64 = 0
	for n != 0 {
		if n == -1 {
			i = -1
			break
		}
		if n%2 != 0 {
			n = n - 3
			i++
		}
		if n != 0 {
			n = n / 2
			i++
		}
	}
	return int(i)
}

func TestNumberofoperations(t *testing.T) {
	var num int64 = 2
	ret := Numberofoperations(num)
	t.Log(ret)
}
