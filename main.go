package main

import (
	"fmt"
	"sync"
)

const (
	Phi     = '\U000003A6'
	美分      = '\U000000A2'
	英镑      = '\U000000A3'
	通用货币符号  = '\U000000A4'
	CNY     = '\U000000A5'
	两条断开的直线 = '\U000000A6'
	节标记     = '\U000000A7'
	著作权符    = '\U000000A9'
	商标      = '\U000000AE'
	度       = '\U000000B0'
	平方      = '\U000000B1'
	立方      = '\U000000B2'
	四分之一    = '\U000000BC'
	二分之一    = '\U000000BD'
	四分之三    = '\U000000BE'
	反问号     = '\U000000BF'
	品牌商品的标记 = '\U00002122'
	品牌服务的标志 = '\U00002120'
)

func submain() {
	fmt.Printf("'%c' '%c' '%c' '%c' '%c'", Phi, 美分, 英镑, 品牌商品的标记, 品牌服务的标志)

}

// 等待所有 goroutine 执行完毕
// 进入死锁
func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})

	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, done, &wg)
	}

	close(done)
	wg.Wait()
	fmt.Println("all done!")
}

func doIt(workerID int, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerID)
	defer wg.Done()
	<-done
	fmt.Printf("[%v] is done\n", workerID)
}
