package getMinStack

import (
	l "ProgrammingGo/log"

	"testing"
)

/*
描述
实现一个特殊功能的栈，在实现栈的基本功能的基础上，再实现返回栈中最小元素的操作。
示例1
输入：
[[1,3],[1,2],[1,1],[3],[2],[3]]
复制
返回值：
[1,2]
复制
备注：
有三种操作种类，op1表示push，op2表示pop，op3表示getMin。你需要返回和op3出现次数一样多的数组，表示每次getMin的答案

1<=操作总数<=1000000
-1000000<=每个操作数<=1000000
数据保证没有不合法的操作

*/

func TestGetMinStack(t *testing.T) {
	in := [][]int{{1, 3}, {1, 2}, {1, 1}, {3}, {2}, {3}}
	ret := getMinStack(in)
	t.Log(ret)
}

/*
1.push 2.pop 3.min
*/

func getMinStack(op [][]int) []int {
	stack := []int{}
	min := 1000000
	ans := []int{}
	for i := 0; i < len(op); i++ {
		if op[i][0] == 2 {
			if len(stack) == 0 {
				l.Debug.Println("没有可以弹出的")
			} else {
				stack = stack[:len(stack)-1]
			}

			for _, v := range stack {
				if v < min {
					min = v
				}
			}
		}
		if op[i][0] == 1 {
			v := op[i][1]
			if v < min {
				min = v
			}
			stack = append(stack, v)
		}

		if op[i][0] == 3 {
			ans = append(ans, min)
		}
	}
	return ans
}
