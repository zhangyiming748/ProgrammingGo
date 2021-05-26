package cow

func gcd(a, b int) int {
	min := minNum(a, b)
	sum := 0
	for i := min; i >0 ; i-- {
		if a%i==0&&b%i==0{
			sum=i
			break
		}
	}
	return sum
}
func minNum(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
