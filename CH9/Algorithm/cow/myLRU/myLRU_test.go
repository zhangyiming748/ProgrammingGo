package myLRU

import (
	"ProgrammingGo/log"
	"testing"
)

/*
设计LRU(最近最少使用)缓存结构，该结构在构造时确定大小，假设大小为K，并有如下两个功能
1. set(key, value)：将记录(key, value)插入该结构
2. get(key)：返回key对应的value值

提示:
1.某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的，然后都会刷新缓存。
2.当缓存的大小超过K时，移除最不经常使用的记录。
3.输入一个二维数组与K，二维数组每一维有2个或者3个数字，第1个数字为opt，第2，3个数字为key，value
   若opt=1，接下来两个整数key, value，表示set(key, value)
   若opt=2，接下来一个整数key，表示get(key)，若key未出现过或已被移除，则返回-1
   对于每个opt=2，输出一个答案
4.为了方便区分缓存里key与value，下面说明的缓存里key用""号包裹

进阶:你是否可以在O(1)的时间复杂度完成set和get操作
*/

func TestMyLRU(t *testing.T) {

	value:=[][]int{{1,1,1},{1,2,2},{1,3,2},{2,1},{1,4,4},{2,2}}
	op:=3
	value2:=[][]int{{1,1,1},{1,2,2},{2,1},{1,3,3},{2,2},{1,4,4},{2,1},{2,3},{2,4}}
	op2:=2
	ret := LRU(value,op)
	ret2:=LRU(value2,op2)
	t.Log(ret,ret2)
}

type Three struct {
	klist []int
	kv    map[int]int
}

func (t *Three) add(k, v int) {
	t.klist = append(t.klist, k)
	t.kv[k] = v
	if len(t.klist) > 3 {
		t.klist = t.klist[1:]
	}
	return
}
func (t *Three) see(k int) int {
	for _, v := range t.klist {
		if v == k {
			t.klist = append(t.klist, k)
			if len(t.klist) > 3 {
				t.klist = t.klist[1:]
			}
			return t.kv[k]
		}
	}
	return -1
}
func LRU(operators [][]int, k int) []int {
	kv := new(Three)
	kv.kv=make(map[int]int)
	kv.klist=[]int{}
	res := []int{}
	for i, v := range operators {
		log.Info.Printf("正在处理第%d组数据\n", i)
		if v[0] == 1 {
			kv.add(v[1], v[2])
		} else if v[0] == 2 {
			ret := kv.see(v[1])
			res = append(res, ret)
		}
	}
	return res
}

//func LRU(operators [][]int, k int) []int {
//	klist:=[]int{}
//	kv:=make(map[int]int)
//	res:=[]int
//	for i, v := range operators {
//		log.Info.Printf("正在处理第%d组数据\n", i)
//		if v[0] == 1 {
//			kv[v[1]]=v[2]
//			klist=append(klist,v[1])
//		}
//		if v[0] == 2 {
//
//		}
//	}
//	return res
//}
