package myChan

import (
	"fmt"
	"time"
	//"sync"
)

func master() {
	ch := make(chan struct{}, 4)
	//var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		output(i, ch)
	}
}
func output(i int, ch chan struct{}) {
	ch <- struct{}{}
	fmt.Printf("这是第%d次输出\n", i)
	time.Sleep(2 * time.Second)
	<-ch
}
