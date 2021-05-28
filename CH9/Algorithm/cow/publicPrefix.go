package cow

import (
	"fmt"
	"math"
	"strings"
)

func publicPrefix(s []string) string {
	short := shorter(s)
	count := 0
	for i := 0; i < short; i++ {
		char := s[count][i]
		for idx, _ := range s {
			if s[idx][count] == char {
				count += 1
			}
		}

	}
	fmt.Println(count)
	ret := s[0][:count]
	return ret
}
func shorter(s []string) int {
	short := math.MaxInt32
	for _, v := range s {
		s := len(v)
		if s < short {
			short = s
		}
	}
	return short
}
func hasprefix(s, p string) bool {
	if strings.HasPrefix(s, p) {
		return true
	} else {
		return false
	}
}
