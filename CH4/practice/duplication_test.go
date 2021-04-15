package practice

import (
	"math/rand"
	"testing"
)

func TestDuplication(t *testing.T) {
	nums := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	newer := duplication(nums)
	t.Log(newer)

}
func TestFlatten(t *testing.T) {
	numss := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	newer := flatten(numss)
	t.Log(newer)
}
func TestGrouplist(t *testing.T) {
	nums := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	ret := make2D(nums, 3)
	t.Log(ret)
}
func BenchmarkGrouplist(b *testing.B) {
	nums := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	for i := 0; i < b.N; i++ {
		b.Log(make2D(nums, rand.Intn(5)+1))
	}
}
