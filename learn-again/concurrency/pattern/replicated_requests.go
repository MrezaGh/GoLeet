package pattern

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork(done <-chan interface{}, id int, wg *sync.WaitGroup, result chan<- int) {

	start := time.Now()
	loadTime := time.Duration(1+rand.Intn(5)) * time.Second
	select {
	case <-done:
	case <-time.After(loadTime):
	}
	select {
	case <-done:
	case result <- id:
	}
	took := time.Since(start)
	if took < loadTime {
		took = loadTime
	}

	fmt.Printf("id:%d took:%v \n", id, took)
	wg.Done()

}

func replicatedRequests() {
	done := make(chan interface{})
	var wg sync.WaitGroup
	result := make(chan int)

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go doWork(done, i, &wg, result)
	}

	received := <-result
	close(done)
	wg.Wait()

	fmt.Printf("id:%v was the fastest", received)
}
