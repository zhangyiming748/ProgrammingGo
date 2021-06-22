package myLRU

import (
	"ProgrammingGo/log"
	"testing"
)

type list struct {
	Length int         //cache容量
	s      []int       //key的切片
	m      map[int]int //key:value
}

func (l *list) push(key, value int) {
	l.m[key] = value
	log.Info.Printf("set %v:%v\n", key, value)
	l.s = append(l.s, key)
	log.Info.Printf("当前的cache列表%v\n", l.s)
	if len(l.s) > l.Length {
		delete(l.m, l.s[0])
		l.s = l.s[1:]
		log.Info.Printf("修剪过的cache列表%v\n", l.s)

	}
	return
}
func (l *list) pop(key int) int {
	if v, ok := l.m[key]; ok {
		value := v
		log.Info.Printf("set %v:%v\n", key, value)
		l.s = append(l.s, key)
		log.Info.Printf("当前的cache列表%v\n", l.s)
		if len(l.s) > l.Length {
			delete(l.m, l.s[0])
			l.s = l.s[1:]
			log.Info.Printf("修剪过的cache列表%v\n", l.s)
		}
		return value
	}
	return -1

}
func LRU(operators [][]int, k int) []int {
	op := operators
	capp := k
	l := new(list)
	l.m = make(map[int]int)
	l.Length = capp
	l.s = make([]int, 0)
	res := make([]int, 0)
	for _, v := range op {
		if len(v) == 3 {
			key := v[1]
			value := v[2]
			l.push(key, value)
		}
		if len(v) == 2 {
			key := v[1]
			value := l.pop(key)
			res = append(res, value)
		}
	}
	return res
}
func TestLRU(t *testing.T) {
	op := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	cap := 3
	ret := LRU(op, cap)
	t.Log(ret)
}
