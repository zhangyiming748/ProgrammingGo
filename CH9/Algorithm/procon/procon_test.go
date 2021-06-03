package procon

import (
	"fmt"
	"sync"
	"testing"
)

func TestProcon(t *testing.T) {
	var wg sync.WaitGroup

	ch := make(chan int, 0)
	wg.Add(1)
	go func() {
		Producter(ch, &wg)
	}()
	wg.Add(1)
	go func() {
		Consumer(ch, &wg)
	}()
	wg.Wait()
}
func Producter(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 30; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}
func Consumer(ch chan int, wg *sync.WaitGroup) {
	for {
		if v, ok := <-ch; ok {
			fmt.Println(v)
		} else {
			break
		}

	}
	wg.Done()
}
