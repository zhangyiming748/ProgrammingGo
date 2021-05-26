package cow

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	ret := fibonacci(n)
	return ret[n-1]
}
func fibonacci(n int) []int {
	num := []int{1, 1}
	for n > 0 {
		num = append(num, num[len(num)-2]+num[len(num)-1])
		n -= 1
	}
	return num
}
