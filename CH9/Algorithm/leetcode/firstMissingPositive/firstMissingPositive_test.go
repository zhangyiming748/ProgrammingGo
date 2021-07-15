package firstMissingPositive

import (
	"math"
	"testing"
)

func firstMissingPositive(nums []int) int {
	m:=make(map[int]bool)
	for _,v:=range nums{
		if v>0{
			m[v]=true
		}
	}
	for i:=1;i<math.MaxInt32;i++{
		if _,ok:=m[i];ok{
			continue
		}else {
			return i
		}
	}
	return nums[len(nums)]+1
}

func TestFirstMissing(t *testing.T) {
	e1 := []int{1, 2, 0}
	//3
	r1 := firstMissingPositive(e1)
	e2 := []int{3, 4, -1, 1}
	//2
	r2 := firstMissingPositive(e2)
	e3 := []int{7, 8, 9, 11, 12}
	r3 := firstMissingPositive(e3)
	//1
	t.Logf("r1=%d\tr2=%d\tr3=%d\n",r1,r2,r3)
}
