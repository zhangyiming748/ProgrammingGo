package cow

import (
	"strconv"
)

func reverse(x int) int {
	var negative bool
	if x < 0 {
		negative = true
		x = 0 - x
	}
	xs := strconv.Itoa(x)
	b := []byte(xs)
	//l.Debug.Printf("b= %v",b)
	nb := []byte{}
	for i := 0; i < len(b); i++ {
		//l.Debug.Printf("当前循环的字母%v",b[len(b)-(i+1)])
		nb = append(nb, b[len(b)-(i+1)])
	}
	ret, _ := strconv.Atoi(string(nb))
	if negative {
		ret = 0 - ret
	}
	return ret
}
