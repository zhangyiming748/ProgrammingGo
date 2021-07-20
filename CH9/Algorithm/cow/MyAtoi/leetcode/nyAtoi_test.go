package leetcode

import "testing"

func TestMyAtoi(t *testing.T) {
	s := "words and 987\n"
	ret := myAtoi(s)
	t.Log(ret)
}
