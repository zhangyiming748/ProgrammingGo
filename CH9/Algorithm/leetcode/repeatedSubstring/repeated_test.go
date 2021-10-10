package repeatedSubstring

import (
	"strings"
	"testing"
)

func repeatedSubstringPattern(s string) bool {
	ns := strings.Join([]string{s, s}, "")
	ns = ns[1 : len(ns)-1]
	if contain := strings.Contains(ns, s); contain {
		return true
	} else {
		return false
	}
}
func TestRepeatedSubstringPattern(t *testing.T) {
	can := "abab"
	cannot := "aba"
	r1 := repeatedSubstringPattern(can)
	r2 := repeatedSubstringPattern(cannot)
	t.Logf("r1=%v,r2=%v\n", r1, r2)
}
