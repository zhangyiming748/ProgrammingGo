package repeatedSubstringPattern

import "testing"

func TestRepeat(t *testing.T) {
	t1 := "abab"
	t2 := "aba"
	ret1 := repeatedSubstringPattern(t1)
	ret2 := repeatedSubstringPattern(t2)
	t.Logf("r1=%v,r2=%v\n", ret1, ret2)
}
