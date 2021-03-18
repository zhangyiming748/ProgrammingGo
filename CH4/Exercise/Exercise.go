package Exercise

func uniqueInts(s []int) []int {
	newSlice := make([]int, 0)
	exist := make(map[int]bool)
	for _, v := range s {
		if _, ok := exist[v]; !ok {
			exist[v] = true
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func flatten(ss [][]int) []int {
	var ret []int
	for _, s := range ss {
		for _, v := range s {
			ret = append(ret, v)
		}
	}
	return ret
}
func make2d(s []int, n int) [][]int {
	var ss [][]int
	//group:=len(s)/n //多少组

	for i := 0; i < len(s); i++ {
		inner:=[]int{}
		for j:=0;j<n;j++{
			inner=append(inner,s[i])
		}
		ss= append(ss, inner)
	}

	return ss
}
