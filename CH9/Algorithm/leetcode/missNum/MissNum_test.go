package missNum

import "testing"

/*集合 s 包含从 1 到 n 的整数
不幸的是，因为数据错误，导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合丢失了一个数字
并且
有一个数字重复
*/
func findErrorNums(nums []int) []int {
	l := len(nums)
	var miss int
	for i := 0; i < l; i++ {
		if i+1 == nums[i] {
			continue
		} else {
			miss = i
			break
		}
	}
	return []int{miss, miss + 1}
}

func TestUnit(t *testing.T) {
	nums1 := []int{1, 2, 2, 4}
	//[2,3]
	nums2 := []int{1, 1}
	//[1,2]
	t.Logf("ret1=%v\tret2=%v\n", findErrorNums(nums1), findErrorNums(nums2))
}
