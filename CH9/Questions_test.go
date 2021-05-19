package CH9

import (
	"fmt"
	"time"

	"testing"
)

func startAll(c1, c2, c3, c4 chan string) {
	go func() {
		m1 := "this is from c1"
		c1 <- m1
	}()
	go func() {
		m2 := "this is from c2"
		c2 <- m2
	}()
	go func() {
		m3 := "this is from c3"
		c3 <- m3
	}()
	go func() {
		m4 := "this is from c1"
		c4 <- m4
	}()

}
func TestSelect(t *testing.T) {
	f1 := make(chan string, 1)
	f2 := make(chan string, 1)
	f3 := make(chan string, 1)
	f4 := make(chan string, 1)

	startAll(f1, f2, f3, f4)
	select {
	case s := <-f1:
		fmt.Println(s)
	case s := <-f2:
		fmt.Println(s)
	case s := <-f3:
		fmt.Println(s)
	case s := <-f4:
		fmt.Println(s)
	default:
		<-time.After(400 * time.Millisecond)

	}
}
