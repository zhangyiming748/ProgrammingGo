package CH9

import (
	"fmt"
	"strings"
)

func Q2() {
	s1 := "fach"
	s2 := "bbaaccedfg"
	ans := ""
	for _, v := range s1 {
		if strings.Contains(s2, string(v)) {
			if !strings.Contains(ans, string(v)) {
				ans += string(v)
			}
		}
	}
	//fmt.Printf("ans = %s \n", ans)
	ret := MySort(ans)
	fmt.Printf("ret = %v\n", ret)

}
func MySort(s string) string {
	r := make([]rune, 0)

	for _, v := range s {
		r = append(r, v)
	}
	for i := 0; i < len(r); i++ {
		for j := 1; j < len(r)-i; j++ {
			if r[j] < r[j-1] {
				//交换
				r[j], r[j-1] = r[j-1], r[j]
			}
		}
	}
	var ret string
	for _, val := range r {
		ret = ret + string(val)
	}
	return ret
}
