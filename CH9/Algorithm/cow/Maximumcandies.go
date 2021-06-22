package cow

//import "log"
func Maximumcandies(n int64, k int64) int64 {
	if k%(n+1) == n {
		return k / (n + 1)
	}
	return k/(n+1) - 1
}
