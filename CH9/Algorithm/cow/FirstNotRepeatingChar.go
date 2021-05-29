package cow

import (
	"fmt"
)

type abc struct {
	index int
}

func FirstNotRepeatingChar(str string) int {
	s := []byte(str)
	fmt.Println(s)

	for i := 0; i <= len(s)-2; i++ {
		for j := i + 1; j <= len(s)-1; j++ {
			if s[i] == s[j] {
				break
			}
			if j == len(s)-1 {
				return i
			}
		}
	}
	return -1
}
