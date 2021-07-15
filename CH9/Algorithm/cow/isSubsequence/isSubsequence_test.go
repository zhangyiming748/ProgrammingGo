package isSubsequence

import (
	"strings"
	"testing"
)

func isSubsequence(s string, t string) bool {
	bs := []byte(s)
	bt := []byte(t)
	l := len(bs)s
	i:=0//index of s
	j:=0//index of t
	for i<=l {
		if bs[i]==bt[j]
	}

}
func TestIsSubsequence(t *testing.T) {
	s := "abc"
	i := "ahbgdc"
	//true
	ret := isSubsequence(s, i)
	t.Log(ret)
}
