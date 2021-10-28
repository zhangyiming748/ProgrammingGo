package binarysearch

import (
	"math/rand"
	"testing"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func TestUnit(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	ret := search(nums, target)
	t.Log(ret)
}
func BenchmarkUnit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{-1, 0, 3, 5, 9, 12}
		target := rand.Intn(12)
		ret := search(nums, target)
		b.Log(ret)
	}
}