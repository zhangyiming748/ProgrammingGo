package HJ22

import (
	"log"
	"testing"
)

func getBottle(n int,total int) int {
	total=n
	for{
		if n<2{}
	}

}
func TestGetBottle(t *testing.T) {
	in := []int{3, 10, 81, 0}
	for i, v := range in {
		log.Printf("round %d\n", i+1)
		ret := getBottle(v, 0)
		log.Println(ret)
	}
}
