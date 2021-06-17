package myLRU

import (
	"ProgrammingGo/log"
	"testing"
)

type nums struct {
	m map[int]int //存kv
	s []int       //存cache的key
}

func (n *nums) get(key int) int {
	for _, v := range n.s {
		if v == key {
			n.cache(v)
			return n.m[v]
		}
	}
	return -1
}
func (n *nums) set(k, v int) {
	n.m[k] = v
	n.cache(k)
}
func (n *nums) cache(i int) {
	n.s = append(n.s, i)
	if len(n.s) > 3 {
		delete(n.m, n.s[0])
		n.s = n.s[1:]
	}
	return
}
func LRU() {
	op := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	//cap := 3
	n := new(nums)
	n.m = make(map[int]int)
	res := []int{}
	for i := 0; i < len(op); i++ {
		log.Info.Printf("开始处理第%d组操作数", i+1)
		if len(op[i]) == 3 {
			k := op[i][1]
			v := op[i][2]
			log.Info.Printf("插入键值对%d:%d\n", k, v)
			n.set(k, v)
		}
		if len(op[i]) == 2 {
			k := op[i][1]

			res = append(res, n.get(k))
			log.Info.Printf("查找键值对%d,此时的res=%v\n", k, res)
		}
	}
	log.Debug.Println()
}
func TestLRU(t *testing.T) {
	//op := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	//cap := 3
	LRU()
}
