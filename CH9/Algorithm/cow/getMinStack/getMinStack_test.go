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
type Stack struct {
	lens []int
	top  int
	min  int
}

func (s Stack) isEmpty() bool {
	if len(s.lens) == 0 {
		return true
	}
	return false
}
func (s *Stack) push(i int) {
	s.lens = append(s.lens, i)

	for _, v := range s.lens {
		if v < s.min {
			s.min = v
		}
	}
	return
}
func (s *Stack) pop() {
	if s.isEmpty() {
		return
	}
	l.Info.Printf("pop前的stack%v\n", s.lens)

	s.lens = s.lens[:len(s.lens)-1]
	l.Info.Printf("pop后的stack%v\n", s.lens)
	for _, v := range s.lens {
		if v < s.min {
			s.min = v
		}
	}
	l.Info.Printf("pop后的stack%v\n", s.lens)

	return

}
func (s *Stack) getMin() int {
	return s.min
}
func (s Stack) getTop() int {
	return s.lens[len(s.lens)-1]
}
func TestStack(t *testing.T) {
	var s Stack
	if s.isEmpty() {
		t.Log("空栈")
	}
	s.push(1)
	if s.isEmpty() {
		t.Log("还是空栈")
	}
	ret := s.getTop()
	t.Log(ret)
	s.push(2)
	t.Log(s.getTop())
	s.pop()
	t.Log(s.getTop())
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
	//var s *Stack
	s := new(Stack)
	s.min = 1000000
	var ans []int
	for i := 0; i < len(op); i++ {
		if op[i][0] == 1 {
			s.push(op[i][1])
		}
		if op[i][0] == 2 {
			s.pop()
		}
		if op[i][0] == 3 {
			if !s.isEmpty() {
				ans = append(ans, s.getMin())
			}
		}
	}
	l.Info.Printf("op[0]`s type is %T", op[0])
	return ans
}
