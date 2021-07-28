package intersect

import "testing"

func intersect(nums1, nums2 []int) []int {
	n1 := duplicate(nums1)
	n2 := duplicate(nums2)
	inner:=[]int{}
	if len(n1) < len(n2) {
		n1, n2 = n2, n1
	}
	for i := 0; i < len(n1); i++ {
		for j:=0;j<len(n2);j++{
			if n1[i]==n2[j]{
				inner=append(inner,n1)
			}
		}
	}
	return inner
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
