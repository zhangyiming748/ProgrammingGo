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

func TestFibonacci(t *testing.T) {
	n := 0
	ret := Fibonacci(n)

	t.Log(ret)
	ret2 := fibonacci(n)
	t.Log(ret2)
}
func TestReString(t *testing.T) {
	s1 := "abcd"
	ret := ReString(s1)
	t.Log(ret)
}
func TestGcd(t *testing.T) {
	a, b := 361, 121
	ret := gcd(a, b)
	t.Log(ret)
}
func TestMount(t *testing.T) {
	nums := []int{2, 4, 1, 2, 7, 8, 4}
	ret := mount(nums)
	t.Log(ret)
}
func TestSpiralOrder(t *testing.T) {
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	ret := spiralOrder(nums)
	t.Log(ret)
}
func TestMoreThanHalfNum_Solution(t *testing.T) {
	n1 := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	ret1 := MoreThanHalfNum_Solution(n1)
	t.Logf("r1=%d\n", ret1)
}
func TestIsValid(t *testing.T) {
	s := "()"
	ret := isValid(s)
	t.Log(ret)
}
func TestExpression_evaluation(t *testing.T) {
	//Expression_evaluation()
}
func TestMerge(t *testing.T) {
	nums1 := []int{2, 3}
	nums2 := []int{1, 4}
	merge(nums1, len(nums1), nums2, len(nums2))
	mergeSort(nums1, nums2)
}
func TestMissNum(t *testing.T) {
	nums1 := []int{0, 1, 2, 3, 4, 5, 7, 8, 9}
	nums2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	ret1 := missNum(nums1)
	ret2 := missNum(nums2)
	t.Logf("ret1 = %d\tret2 = %d\n", ret1, ret2)
}
func TestPublicPrefix(t *testing.T) {
	words := []string{"abc", "abc", "abca", "abc", "abcc"}
	ret := publicPrefix(words)
	//ret:=shorter(words)
	t.Log(ret)
}
