package mathexp

import "testing"

func TestMathexp(t *testing.T) {
	var n1 int64 =55
	var n2 int64 =10
	ret1:=mathexp(n1)
	ret2:=mathexp(n2)
	t.Log(ret1,ret2)
}
