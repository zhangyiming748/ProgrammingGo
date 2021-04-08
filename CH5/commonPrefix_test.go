package CH5

import (
	"fmt"
	"testing"
)

func TestCommonPathPrefix(t *testing.T) {
	testData := [][]string{
		{"/home/user/goeg", "/home/user/goeg/prefix",
			"/home/user/goeg/prefix/extra"},
		{"/home/user/goeg", "/home/user/goeg/prefix",
			"/home/user/prefix/extra"},
		{"/pecan/π/goeg", "/pecan/π/goeg/prefix",
			"/pecan/π/prefix/extra"},
		{"/pecan/π/circle", "/pecan/π/circle/prefix",
			"/pecan/π/circle/prefix/extra"},
		{"/home/user/goeg", "/home/users/goeg",
			"/home/userspace/goeg"},
		{"/home/user/goeg", "/tmp/user", "/var/log"},
		{"/home/mark/goeg", "/home/user/goeg"},
		{"home/user/goeg", "/tmp/user", "/var/log"},
	}
	for _, data := range testData {
		fmt.Printf("[")
		gap := ""
		for _, datum := range data {
			fmt.Printf("%s\"%s\"", gap, datum)
			gap = " "
		}
		fmt.Println("]")
		cp := CommonPrefix(data)
		cpp := CommonPathPrefix(data)
		equal := "=="
		if cpp != cp {
			equal = "!="
		}
		fmt.Printf("char ⨉ path prefix: \"%s\" %s \"%s\"\n\n",
			cp, equal, cpp)
	}
}
