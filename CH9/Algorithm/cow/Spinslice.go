package cow

import "log"

func spinSlice(n, m int, a []int) []int {
	return solve(n, m, a)
}
func solve(n int, m int, a []int) []int {
	l := n - 1
	for i := 0; i < m; i++ {
		tmp := a[l:]
		a = a[:l]
		a = append(tmp, a...)
	}
	log.Println(a)
	return a
}
