package cow

import (
	"fmt"
	"sort"
)

func merge(A []int, m int, B []int, n int) {
	if len(A) == 0 {
		if len(B) == 0 {
			return
		} else {
			A = B
			fmt.Println(A)
			return
		}
	}
	nline := make([]int, 0)
	var i, j int = 0, 0
	for i < m-1 && j < n-1 {
		if A[i] > B[j] {
			nline = append(nline, B[j])
			j += 1
		}
		if A[i] <= B[j] {
			nline = append(nline, A[i])
			i += 1
		}
	}
	if i == m-1 {
		nline = append(nline, A[m-1])
		nline = append(nline, B[j:]...)
	}
	if j == n-1 {
		nline = append(nline, B[n-1])
		nline = append(nline, A[i:]...)
	}
	A = nline
	fmt.Println(A)
	return
}
func mergeSort(A, B []int) {
	A = append(A, B...)
	fmt.Printf("before: %v", A)
	sort.Ints(A)
	fmt.Printf("after: %v", A)
}
