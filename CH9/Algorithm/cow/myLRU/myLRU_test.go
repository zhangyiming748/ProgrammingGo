package myLRU

import (
	"log"
	"testing"
)

type myMaps struct {
	limit int
	keys  []int
	value []int
}

func (m *myMaps) get(n int) int {
	for i := len(m.keys); i > len(m.keys)-m.limit; i-- {
		if m.keys[i] == n {
			m.keys = append(m.keys, i)
			m.value = append(m.value, m.value[i])
			return m.value[i]
		}
	}
	return -1
}
func (m *myMaps) set(k, v int) {
	m.keys = append(m.keys, k)
	m.value = append(m.value, v)
	return
}
func TestMyLRU(t *testing.T) {
	op := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	k := 3
	ret := LRU(op, k)
	t.Log(ret)
}

func LRU(operators [][]int, k int) []int {
	var lru myMaps
	lru.limit = k
	res:=[]int{}
	for i, v := range operators {
		log.Printf("round %d\n", i+1)
		if v[0] == 1 {
			k:=v[1]
			v:=v[2]
			lru.set(k,v)
		} else if v[0] == 2 {
			k:=v[1]
			res=append(res,lru.get(k))
		} else {
			log.Println("wrong op")
			continue
		}
	}
	return res
}
