package MyBench

import (
	"fmt"
	"sync"
)

func Producer(ch chan int,wg *sync.WaitGroup){
	for i:=0;i<10;i++{
		wg.Add(1)
		ch <- i
	}
}
func Consumer(ch chan int, wg *sync.WaitGroup)  {
	for {
		if v,ok:=<-ch;ok{
			wg.Done()
			fmt.Println(v)
		}else {
			break
		}
	}
}
