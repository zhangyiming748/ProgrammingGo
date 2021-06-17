package myLRU

import (
	"ProgrammingGo/log"
	"testing"
)


type nums struct {
	m map[int]int //存kv
	s []int //存cache的key
}

func LRU() {
	op := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	cap := 3
	n := new(nums)
	n.m:=make(map[int]int)
	for i := 0; i < len(op); i++ {
		log.Info.Printf("开始处理第%d组操作数", i+1)
		if len(op[i])==3{
			k:=op[i][1]
			v:=

		}
	}
}
func TestLRU(t *testing.T) {
	op := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	cap := 3
}
