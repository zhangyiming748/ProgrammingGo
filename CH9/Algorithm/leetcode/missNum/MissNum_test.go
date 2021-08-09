package missNum

import "testing"

/*集合 s 包含从 1 到 n 的整数
不幸的是，因为数据错误，导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合丢失了一个数字
并且
有一个数字重复
*/
func findErrorNums(nums []int) []int {
	if nums[0] != 1 {
		return []int{2, 1}
	}
	if nums[1] == 1 {
		return []int{1, 2}
	}
	for i := 0; i < len(nums); i++ {
		if i+1 == nums[i] {
			continue
		} else {
			if i == len(nums)-1 {
				return []int{nums[i], nums[i] - 2}
			} else {
				return []int{nums[i], nums[i] + 1}
			}
		}
	}
	return nil
}

func TestUnit(t *testing.T) {
	nums1 := []int{1, 2, 2, 4}
	//[2,3]
	nums2 := []int{1, 1}
	//[1,2]
	nums3 := []int{2, 2}
	//[2,1]
	t.Logf("ret1=%v\tret2=%v\tret3=%v\n", findErrorNums(nums1), findErrorNums(nums2), findErrorNums(nums3))
}
