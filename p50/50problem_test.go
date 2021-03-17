package p50

import (
	"fmt"
	"testing"
)

func TestParantheses(t *testing.T) {
	println("hello world")
}

type info struct {
	result string
}
func Test()  {
	var data info
	var err error    // err 需要预声明

	data.result, err = work()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("info: %+v\n", data)
}