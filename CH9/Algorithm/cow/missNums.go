package cow

func missNum(a []int) int {
	l := len(a)
	for i := 0; i < l; i++ {
		if a[i] == i {
			continue
		} else {
			return i
		}
	}
	return l
}
