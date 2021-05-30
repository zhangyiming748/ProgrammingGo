package cow

import (
	"fmt"
)

func merge(A []int, m int, B []int, n int) {
	// write code here
	ms := []int{}
	for _, v := range A {
		ms = append(ms, v)
	}
	for _, val := range B {
		ms = append(ms, val)
	}
	fmt.Println(ms)
}
