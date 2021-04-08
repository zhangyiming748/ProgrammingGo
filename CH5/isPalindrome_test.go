package CH5

import "testing"

func TestIsPalindrome(t *testing.T) {
	words:=[]string{"level","age","i"}
	for i,v:=range words{
		t.Log(i,v,isPalindrome(v))
	}

}
