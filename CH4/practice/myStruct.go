package practice

type numbers struct {
	a float64
	b float64
}

func multi(nums numbers) float64 {
	result := nums.a * nums.b
	return result
}
func (n *numbers) setnumber(a, b float64) {
	n.a = a
	n.b = b
	return
}
