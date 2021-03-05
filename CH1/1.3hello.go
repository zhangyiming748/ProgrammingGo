package CH1

import (
	"fmt"
	"os"
	"strings"
)

func Hello() {
	who := "world"
	if len(os.Args) < 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("hello", who)
}
