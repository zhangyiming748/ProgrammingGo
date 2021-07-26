package isSubsequence

import (

	"testing"
)

func isSubsequence(s string, t string) bool {
	bs:=[]byte(s)
	bt:=[]byte(t)
	i,j:=0,0
	for i<len(bs){
		if bs[i]==bt[j] {
			i++
			j++
		}
		if bs[i]!=bt[j]{
			j++
		}
	}

}
func TestIsSubsequence(t *testing.T) {
	s := "abc"
	i := "ahbgdc"
	//true
	ret := isSubsequence(s, i)
	t.Log(ret)
}
