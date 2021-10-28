package selfDividingNumbers

import (
	"ProgrammingGo/log"
	"testing"
)

func selfDiv(n int) bool {
	num := n
	for n > 0 {
		if n == 0 || n%10 == 0 {
			return false
		}
		a := n % 10
		//log.Info.Println(n)
		//log.Info.Println(a)

		if num%a != 0 {
			return false
		}

		n = n / 10

	}
	return true
}

func selfDividingNumbers(left int, right int) []int {
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13,14,15}
	nums := []int{}
	for i := left; i <= right; i++ {
		nums = append(nums, i)
	}
	log.Info.Println(nums)
	enable := []int{}
	log.Info.Println(nums)
	for _, v := range nums {
		if selfDiv(v) {
			enable = append(enable, v)
		}
	}
	return enable
}
func TestSelfDividingNumbers(t *testing.T) {
	i := 1
	j := 22
	ret := selfDividingNumbers(i, j)
	t.Log(ret)
}
func TestSelfDiv(t *testing.T) {
	n := 21
	ret := selfDiv(n)
	t.Log(ret)
}
