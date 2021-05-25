package cow

import "testing"

//两数之和
func TestTwoSum(t *testing.T) {
	s := []int{3, 2, 4}
	target := 6
	ret := twoSum(s, target)
	t.Log(ret)
}

//反转链表
func TestReverseList(t *testing.T) {
	l := new(ListNode)
	l.Val = 1
	l1 := new(ListNode)
	l.Next = l1
	l1.Val = 2
	l2 := new(ListNode)
	l1.Next = l2
	l2.Val = 3
	t.Logf("l.val=%d,next=%v\nl1.val=%d,next=%v\nl2.val=%d,next=%v", l.Val, l.Next, l1.Val, l1.Next, l2.Val, l2.Next)
	ret := ReverseList(l)
	ret1 := ret.Next
	ret2 := ret1.Next
	t.Logf("ret.val=%d,ret.next=%v\nret1.val=%d,ret1.next=%v\nret2.val=%d,ret2next=%v\n", ret.Val, ret.Next, ret1.Val, ret1.Next, ret2.Val, ret2.Next)
}
func TestGetLeastNumbers_Solution(t *testing.T) {
	nums := []int{4, 5, 6, 2, 7, 3, 8, 1}
	num := 4
	ret := GetLeastNumbers_Solution(nums, num)
	t.Log(ret)
}
func TestFindMin(t *testing.T) {
	nums := []int{4, 5, 1, 6, 2, 7, 3, 8}
	v := findMin(nums)
	t.Logf("i=%d\tv=%d", v.index, v.value)
}
func TestFindMax(t *testing.T) {
	//nums := []int{4, 5, 1, 6, 2, 7, 3, 8}

}
func TestDelMax(t *testing.T) {
	nums := []int{4, 5, 1, 6, 2, 7, 3, 8}
	for i := 0; i < 4; i++ {
		newn := delMax(nums)
		t.Log(newn)
		nums = newn
	}
}
func TestDelMin(t *testing.T) {
	nums := []int{4, 5, 1, 6, 2, 7, 3, 8}
	delMin(nums)

}
func TestProcess(t *testing.T) {
	nums := []int{4, 5, 1, 6, 2, 7, 3, 8}
	newer := append(nums[:3], nums[4:]...)
	t.Log(newer)
}
