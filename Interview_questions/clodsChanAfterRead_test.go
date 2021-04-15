package Interview_questions

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

/**
go channel close后读的问题
可以读取
写入会导致panic
*/

func producer(c chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	wg.Done()
}
func customer(c chan int, wg *sync.WaitGroup) {
	for {
		if v, ok := <-c; ok {
			log.Println(v)
		} else {
			break
		}
	}
	wg.Done()
}
func TestMaster(t *testing.T) {
	c := make(chan int, 8)
	var wg sync.WaitGroup
	wg.Add(1)
	go producer(c, &wg)
	wg.Add(1)
	go customer(c, &wg)
	wg.Wait()
}
func TestAnother(t *testing.T) {
	defer func() {

		if err := recover(); err != nil {
			t.Log(err)
		}
	}()
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for value := range ch {
		fmt.Println("value:", value)
	}
	i := 3
	ch <- i
}
