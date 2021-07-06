package myLRU

import (
	"log"
	"testing"
)

type myMaps struct {
	limit int
	keys  []int
	kv    map[int]int
}

func (m *myMaps) set(k, v int) {
	m.kv[k] = v
	m.keys = append(m.keys, k)
}
func (m *myMaps) get(n int) int {
	if len(m.keys) < 3 {
		for _, v := range m.keys {
			if n == v {
				m.keys = append(m.keys, n)
				return m.kv[n]
			}
		}
		return -1
	}
	//len>=3
	for i := len(m.keys); i > len(m.keys)-m.limit; i-- {
		if m.keys[i] == n {
			m.keys = append(m.keys, n)
			return m.kv[n]
		}
	}
	return -1

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
	lru.kv=make(map[int]int)
	res := []int{}
	for i, v := range operators {
		log.Printf("round %d\n", i+1)
		if v[0] == 1 {
			k := v[1]
			v := v[2]
			lru.set(k, v)
		} else if v[0] == 2 {
			k := v[1]
			res = append(res, lru.get(k))
		} else {
			log.Println("wrong op")
			continue
		}
	}
	return res
}
