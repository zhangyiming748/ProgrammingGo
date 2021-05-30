package cow

import (
	"sort"
)

func findKth(nums []int, n, k int) int {
	dump := dumplicate(nums)
	n = len(dump)
	sort.Ints(dump)
	return dump[n-k]
}
func dumplicate(nums []int) []int {
	m := make(map[int]bool)
	dump := []int{}
	for _, v := range nums {
		if m[v] {
			continue
		} else {
			m[v] = true
			dump = append(dump, v)
		}
	}
	return dump
}
