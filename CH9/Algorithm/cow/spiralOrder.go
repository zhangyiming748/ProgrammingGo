package cow

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	length := len(matrix)
	if length == 0 {
		return []int{}
	}
	var (
		top    = 0
		bottom = length - 1
		right  = len(matrix[0]) - 1
		left   = 0
	)
	for top < bottom && left < right {
		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}
		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right])
		}
		for i := right; i > left; i-- {
			res = append(res, matrix[bottom][i])
		}
		for i := bottom; i > top; i-- {
			res = append(res, matrix[i][left])
		}
		right -= 1
		left += 1
		top += 1
		bottom -= 1
	}
	if bottom == top {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
	} else if left == right {
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][left])
		}
	}
	return res
}
