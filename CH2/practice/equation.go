package CH2

func equation(a, b, c float64) (y1, y2 float64) {

	if θ, ok := isEffective(a, b, c); ok {
		y1 = ((0 - b) + θ) / 2 * a
		y2 = ((0 - b) - θ) / 2 * a
		return y1, y2
	}
	return -1, -1
}
func isEffective(a, b, c float64) (ret float64, able bool) {
	if ret := (b * b) - (4 * a * c); ret >= 0 {
		return ret, true
	}
	return -1.0, false

}
