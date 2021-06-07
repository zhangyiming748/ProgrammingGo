package multiBig

func solve(a []int) int64 {
	if len(a) == 3 {
		return int64(a[0] * a[1] * a[2])
	}
	var z1, z2, z3, b1, b2, sum1, sum2 int64
	var i1, i2, j1 int
	for i := 0; i < len(a); i++ {
		if int64(a[i]) > z1 {
			z1 = int64(a[i])
			i1 = i
		}
		if int64(a[i]) < b1 {
			b1 = int64(a[i])
			j1 = i
		}
	}
	for i := 0; i < len(a); i++ {
		if int64(a[i]) > z2 && i != i1 {
			z2 = int64(a[i])
			i2 = i
		}
		if int64(a[i]) < b1 && i != j1 {
			b2 = int64(a[i])
		}
	}
	for i := 0; i < len(a); i++ {
		if int64(a[i]) > z3 && i != i1 && i != i2 {
			z3 = int64(a[i])
		}
	}
	sum1 = z1 * z2 * z3
	sum2 = z1 * b1 * b2
	if sum1 > sum2 {
		return sum1
	} else {
		return sum2
	}
}
