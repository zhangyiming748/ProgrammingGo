package CH9

import "fmt"

/*
标题：按身高和体重排队 时间限制：1秒 内存限制：262144K 语言：不限

某学校举行运动会，学生们按编号1、2、3、....n进行标识，现需要按照身高由低到高排列，对身高相同的人，按体重由轻到重排列；对身高和体重都相同的人，维持原有的编号顺序关系。请输出排列后的学生编号。

输入描述：两个序列，每个序列由n个正整数组成0<n<=100。第一个序列中的数值代表身高，第二个序列中的数值代表体重。
输出描述：排列结果，每个数值都是原始序列中的学生编号，编号从1开始

比如：
输入：4 100 100 120 130 40 30 60 50
输出：2 1 3 4
*/
type Student struct {
	Height int
	Weight int
}

func (s *Student) SetH(i int) {
	s.Height = i
	return
}
func (s *Student) SetW(i int) {
	s.Weight = i
	return
}
func (s Student) GetH() int {
	return s.Height
}
func (s Student) GetW() int {
	return s.Weight
}
func Q1() {

	list := []int{4, 100, 100, 120, 130, 40, 30, 60, 50}
	m := make(map[int]Student)
	for i := 0; i < list[0]; i++ {
		h:=list[i+1]
		w:=list[i+5]
		m[i+1]=setHW(h,w)
	}

	for key,value:=range m{
		fmt.Println(key,value)
	}
}
func setHW(h,w int)(s Student){
	s.SetH(h)
	s.SetW(w)
	return s
}
func desc(m map[int]Student) {

	for k,v:=range m{

	}
}
