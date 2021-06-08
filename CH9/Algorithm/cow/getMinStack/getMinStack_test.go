package getMinStack

import (
	l "ProgrammingGo/log"

	"math"
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
type Stack struct {
	nums []int
	min  int
}

func (s *Stack) pop() int {
	val := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	for _, v := range s.nums {
		if v < s.min {
			s.min = v
		}
	}
	return val
}
func (s *Stack) push(i int) {
	s.nums = append(s.nums, i)
	for _, v := range s.nums {
		if v < s.min {
			s.min = v
		}
	}
	return
}

func TestStack(t *testing.T) {
	s := new(Stack)
	s.push(1)
	s.push(2)

}
func TestGetMinStack(t *testing.T) {
	in := [][]int{{1, 3}, {1, 2}, {1, 1}, {3}, {2}, {3}}
	ret := getMinStack(in)
	t.Log(ret)
}

/*
1.push 2.pop 3.min
*/

func getMinStack(op [][]int) []int {
	s := new(Stack)
	s.min = math.MaxInt32
	ans := []int{}
	for i := 0; i < len(op); i++ {
		for j := 0; j < len(op[i]); j++ {
			l.Info.Printf("一次循环%v\n", op[i][j])
			if op[i][j] == 1 && j == 0 {
				s.push(op[i][j+1])
			}
			if op[i][j] == 2 && j == 0 {
				s.pop()
			}
			if op[i][j] == 3 && j == 0 {
				ans = append(ans, s.min)
			}
		}

		//if op[i][0] == 3 {
		//	ans = append(ans, s.min)
		//}
		//if op[i][0] == 2 {
		//	p:=s.pop()
		//	l.Info.Println("弹出",p)
		//}
		//if op[i][0] == 1 {
		//	s.push(op[i][1])
		//}
	}
	return ans
}
