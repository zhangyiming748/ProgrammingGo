package aboutStrings

import "testing"

func TestReplace(t *testing.T) {
	input1 := "GeeksforGeeks is a computer science portal."
	output1 := replace(input1)
	t.Logf("out 1 = %v", output1)
	input2 := "Hello, Welcome to GeeksforGeeks"
	output2 := replace(input2)
	t.Logf("out 2 = %v", output2)
}
