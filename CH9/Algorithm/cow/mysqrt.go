package cow

func mySqrt(x int) int {
	count := 1
	for count < x/2 {
		if x > count*count {
			count++
		}
		if x == count*count {
			return count
		}
		if x < count*count {
			break
		}
	}
	if count*count > x {
		count--
	}
	return count
}
