package practice



func duplication(i []int) []int {
	m := make(map[int]bool)
	var s []int
	for _, v := range i {
		if _, ok := m[v]; !ok {
			m[v] = true
			s = append(s, v)
		}
	}
	return s
}
func flatten(ss[][]int)[]int {
	nums:=[]int{}
	for _,s:=range ss {
		for _,v:=range s{
			nums= append(nums, v)
		}
	}
	return nums
}
func make2D(slice []int, columns int) [][]int {
	matrix := make([][]int, neededRows(slice, columns))
	for i, x := range slice {
		row := i / columns
		column := i % columns
		if matrix[row] == nil {
			matrix[row] = make([]int, columns)
		}
		matrix[row][column] = x
	}
	return matrix
}

func neededRows(slice []int, columns int) int {
	rows := len(slice) / columns
	if len(slice)%columns != 0 {
		rows++
	}
	return rows
}