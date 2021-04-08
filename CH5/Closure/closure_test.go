package Closure

import "testing"

func TestClosure(t *testing.T) {
	closure()
}
func TestFactoryfunc(t *testing.T) {
	addZip:=factoryfunc(".zip")
	t.Log(addZip("filename"))
}
