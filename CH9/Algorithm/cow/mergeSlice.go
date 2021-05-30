package cow

import (
	"fmt"
	"sort"
)

func merge(A []int, m int, B []int, n int) {
	// write code here
	ms := []int{}
	if m == 0 && n != 0 {
		sort.Ints(B)
		A = B
	}
	if m != 0 && n == 0 {
		sort.Ints(A)
	}
	if m == 0 && n == 0 {

	}
	for _, v := range A {
		ms = append(ms, v)
	}
	for _, val := range B {
		ms = append(ms, val)
	}
	sort.Ints(ms)
	fmt.Println(ms)
	A = ms
}
