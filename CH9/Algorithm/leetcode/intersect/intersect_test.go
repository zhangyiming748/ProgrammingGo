package intersect

import (
	"sort"
	"testing"
)

func intersect(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	ret:=make([]int,0)
	i,j:=0,0
	for i<len(nums1)||j<len(nums2){
		if nums1[i]<nums2[j]{
			i++
		}else if nums1[i]>nums2[j]{
			j++
		}else {
			ret=append(ret,nums1[i])
			i++
			j++
		}
	}
	return ret
}
func duplicate(n []int) []int {
	newnums := []int{}
	kv := make(map[int]bool)
	for _, v := range n {
		if _, ok := kv[v]; !ok {
			kv[v] = true
			newnums = append(newnums, v)
		}
	}
	return newnums
}
func TestUnit(t *testing.T) {
	n := []int{1, 1, 2, 3, 4, 4, 4, 5, 4}
	ret := duplicate(n)
	t.Log(ret)
}
