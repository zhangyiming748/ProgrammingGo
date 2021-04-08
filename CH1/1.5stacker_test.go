package CH1

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	var haystack Stack
	haystack.Push("hey")
	haystack.Push(15)
	haystack.Push([]string{"i`m", "good", "thanks"})
	haystack.Push(81.52)
	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
