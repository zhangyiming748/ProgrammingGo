package procon

//func Producter(ch chan int,wg *sync.WaitGroup) {
//	for i := 0; i < 30; i++ {
//		ch <-i
//	}
//	close(ch)
//	wg.Done()
//}
//func Consumer(ch chan int,wg *sync.WaitGroup)  {
//	for {
//		if v,ok:=<-ch;ok{
//			fmt.Println(v)
//		}else {
//			break
//		}
//
//	}
//	wg.Done()
//}
