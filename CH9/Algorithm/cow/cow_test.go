package cow

import (
	"fmt"
	"strconv"
	"testing"
)

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

type slice []int

func TestGo(t *testing.T) {
	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Print(v)
	}

}

func (s *slice) Add(exem int) *slice {
	fmt.Println(exem)
	*s = append(*s, exem)
	return s
}

func TestFirstNotRepeatingChar(t *testing.T) {
	str := "google"
	ret := FirstNotRepeatingChar(str)
	t.Log(ret)
}
func TestMysqrt(t *testing.T) {
	in := 7
	ret := mySqrt(in)
	t.Log(ret)
}

func TestFindKth(t *testing.T) {
	nums := []int{1, 3, 5, 2, 2}
	//1,2,3,5 3
	ret1 := dumplicate(nums)
	t.Logf("ret1 = %v", ret1)
	l := 5
	target := 3
	ret := findKth(nums, l, target)
	t.Log(ret)
}
func TestByte(t *testing.T) {
	s := "0123456789ABCDEFGabcdefg"
	b := []byte(s)
	for i, v := range b {
		t.Logf("%d `s ascii is %v\n", i+1, v)
	}
}
func TestEffectiveIp(t *testing.T) {
	ipv4 := "172.16.254.1"
	ipv6 := "2001:0db8:85a3:0:0:8A2E:0370:7334"
	t1 := EffectiveIp(ipv4)
	t.Logf("ipv4 is %s", t1)
	t2 := EffectiveIp(ipv6)
	t.Logf("ipv6 is %s", t2)
}
func TestReverse(t *testing.T) {
	in1 := 123
	//49 50 51
	in2 := -123
	ret1 := reverse(in1)
	ret2 := reverse(in2)
	t.Logf("ret1 = %d\tret2 = %d\n", ret1, ret2)
}
func TestMaxLength(t *testing.T) {
	nums := []int{2, 2, 3, 4, 8, 99, 3}
	ret := maxLength(nums)
	t.Log(ret)
}
func TestName(t *testing.T) {

	var s string
	fmt.Scanln(&s)
	h, _ := strconv.ParseFloat(s, 64)
	long := (h - 105) * 0.618 / 3.14
	fmt.Printf("长度= %vcm", long)

}
func TestLongestCommonPrefix(t *testing.T) {
	str := []string{"abca", "abc", "abca", "abc", "abcc"}
	//ret1 := getShort(str)
	ret := longestCommonPrefix(str)
	t.Log(ret)
}
