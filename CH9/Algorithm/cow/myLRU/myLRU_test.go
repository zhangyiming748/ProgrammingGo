package myLRU

import (
	"ProgrammingGo/log"
	"testing"
)

/*
描述
设计LRU缓存结构，该结构在构造时确定大小，假设大小为K，并有如下两个功能
set(key, value)：将记录(key, value)插入该结构
get(key)：返回key对应的value值
[要求]
set和get方法的时间复杂度为O(1)
某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的。
当缓存的大小超过K时，移除最不经常使用的记录，即set或get最久远的。
若opt=1，接下来两个整数x, y，表示set(x, y)
若opt=2，接下来一个整数x，表示get(x)，若x未出现过或已被移除，则返回-1
对于每个操作2，输出一个答案
输入：[[1,1,1],[1,2,2],[1,3,2],[2,1],[1,4,4],[2,2]],3
返回值：[1,-1]
第一次操作后：最常使用的记录为("1", 1)
第二次操作后：最常使用的记录为("2", 2)，("1", 1)变为最不常用的
第三次操作后：最常使用的记录为("3", 2)，("1", 1)还是最不常用的
第四次操作后：最常用的记录为("1", 1)，("2", 2)变为最不常用的
第五次操作后：大小超过了3，所以移除此时最不常使用的记录("2", 2)，加入记录("4", 4)，并且为最常使用的记录，然后("3", 2)变为最不常使用的记录
*/
type list struct {
	nums  map[int]int
	cache []int //缓存列表
	cap   int   //缓存容量
}

func (l *list) set(k, v int) {
	l.nums[k] = v
	l.cache = append(l.cache, k)
	if len(l.cache) > l.cap {
		l.cache = l.cache[1:]
	}
	return
}
func (l *list) get(k int) int {
	for _, v := range l.cache {
		if v == k {
			l.cache = append(l.cache, k)
			if len(l.cache) > l.cap {
				l.cache = l.cache[1:]
				return l.nums[v]
			}
		}
	}
	return -1
}

func TestLRU(t *testing.T) {
	in := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	cache := 3
	ret := LRU(in, cache)
	newret := allIn(in, cache)
	t.Log(ret)
	t.Log(newret)
}
func LRU(operators [][]int, k int) []int {
	// write code here
	var res []int
	li := new(list)
	li.nums = make(map[int]int)
	li.cap = k
	li.cache = make([]int, 0)
	for i, v := range operators {
		if v[0] == 1 {
			key := v[1]
			val := v[2]
			li.set(key, val)
			log.Info.Printf("第%d组set操作:key=%d\tval=%d\n", i+1, key, val)
			log.Info.Printf("此时的map=%v\ncache列表=%v\n容量设置为%v\n", li.nums, li.cache, li.cap)
		} else if v[0] == 2 {
			key := v[1]
			ret := li.get(key)
			res = append(res, ret)
			log.Info.Printf("第%d组get操作:key=%d\tval=%d\n", i+1, key, ret)
			log.Info.Printf("此时的map=%v\ncache列表=%v\n容量设置为%v\n", li.nums, li.cache, li.cap)
		}
	}
	return res
}
func allIn(op [][]int, k int) []int {
	cache := make([]int, 0) //缓存队列
	kv := make(map[int]int) //保存所有键值对
	res := make([]int, 0)
	for i, v := range op {
		log.Info.Printf("循环第%d组操作数", i+1)
		if v[0] == 1 && len(v) == 3 { //set操作
			key := v[1]
			val := v[2]
			kv[key] = val
			cache = append(cache, key)
			if len(cache) > k {
				cache = cache[1:]
			}
		}
		if v[0] == 2 && len(v) == 2 {
			for i, val := range cache {
				if val == v[1] {
					res = append(res, kv[val])
					cache = append(cache, val)
					if len(cache) > k {
						cache = cache[1:]
					}
				}
				if i == len(cache)-1 && val != v[1] {
					res = append(res, -1)
				}
			}
		}
	}
	return res
}
