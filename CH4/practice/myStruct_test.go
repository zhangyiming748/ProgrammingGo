package practice

import "testing"

func TestMyStruct(t *testing.T) {
	var (
		nums numbers
	)
	nums.setnumber(3.0, 4.5)
	result := multi(nums)
	t.Logf("%f", result)
}
