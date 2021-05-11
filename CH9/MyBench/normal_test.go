package MyBench

import (
	"sync"
	"testing"
)

func TestModel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int, 1)
	Producer(ch, &wg)

	Consumer(ch, &wg)
	wg.Wait()
}
func TestProducer(t *testing.T) {

}
