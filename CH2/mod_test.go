package CH2

import "testing"

func TestMode(t *testing.T) {
	nums := []float64{3.2, 3.5, 3.5, 3.1, 3.2, 3.5}
	ret := mode(nums)
	t.Log(ret)
}
