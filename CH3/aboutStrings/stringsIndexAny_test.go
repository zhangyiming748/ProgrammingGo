package aboutStrings

import (
	//"strings"
	"testing"
)

func TestReturnFirst(t *testing.T) {
	strs:=[...]string{"GeeksforGeeks - A Computer Science Portal","GFG is the Best"}
	for i,v:=range strs{

		t.Logf("%d`s string is %s\n",i,v)
		ret:=returnFirst(v, "A")
		t.Logf("A第一次出现的索引是%d",ret)
	}
}
