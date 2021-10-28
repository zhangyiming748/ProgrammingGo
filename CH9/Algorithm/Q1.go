package Algorithm

import (
	"fmt"
	"sort"
)

type Student struct {
	No     int
	Height int
	Weight int
}

func (s *Student) SetNO(i int) {
	s.No = i
	return
}
func (s Student) GetNO() int {
	return s.No
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
func Q1(input []int) []int {
	student := make([]*Student, 0, input[0])
	for i := 0; i < input[0]; i++ {
		s := &Student{}
		s.SetNO(i + 1)
		s.SetH(input[i+1])
		s.SetW(input[i+5])
		student = append(student, s)
	}
	fmt.Println("before sort")
	for _, v := range student {
		fmt.Println(*v)
	}
	sort.Slice(student, func(i, j int) bool {
		if student[i].Height == student[j].Height {
			return student[i].Weight < student[j].Weight
		}
		return student[i].Height < student[j].Height
	})
	fmt.Println("after sort")
	for _, v := range student {
		fmt.Println(*v)
	}
	ret := []int{}
	for _, v := range student {
		ret = append(ret, v.No)
	}
	return ret

}
