package CH9

import (
	"fmt"
	"testing"
	"time"
)

func TestQ3(t *testing.T) {
	inputs := [...]int{1324567980, 68879077, 89776677, 2615371}
	expected := [...]int{132450, 6077, 6677, 211}

	for i := 0; i < len(inputs); i++ {
		ret := FindMin(inputs[i], 4)
		if ret != expected[i] {
			t.Errorf("input is %d\texpected is %d\tbut actual is %d\n", inputs[i], expected[i], ret)
		}
	}
}
func BenchmarkQ3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindMin(1324567980,4)
	}
	b.StopTimer()
}

func TestBoss(t *testing.T) {
	arr:=[]int{1,2,3}
	for _,a:=range arr{
		go func() {
			time.Sleep(10*time.Millisecond)
			if a>2{
				defer fmt.Println("a>2")
			}else {
				defer fmt.Println("a<2")
			}
		}()
	}
	fmt.Println("end")
	time.Sleep(10*time.Second)
}
func adder()func(int)int{
	sum:=0
	return func(x int) int {
		sum+=x
		return sum
	}
}

