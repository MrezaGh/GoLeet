package main

import (
	"fmt"
	"sync"
	"time"
)

func fanIn(workerChans ...<-chan int) <-chan int {
	results := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(workerChans))

	for _, workerChan := range workerChans {
		go func(workerChan <-chan int) {
			for result := range workerChan {
				results <- result
			}
			wg.Done()
		}(workerChan)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results

}

func fanOut(workerNum int, jobs <-chan int) []<-chan int {
	workerResult := make([]<-chan int, workerNum)
	for i := 1; i <= workerNum; i++ {
		workerResult[i-1] = worker(i, jobs)
	}

	return workerResult
}

func worker(id int, jobs <-chan int) <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)
		for job := range jobs {
			fmt.Println(job, " doing the job")
			time.Sleep(time.Duration(job) * time.Millisecond * 100)
			fmt.Println(job, " did it's job")
			result <- job
		}
	}()

	return result
}

func main() {
	jobs := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	workerResult := fanOut(3, jobs)
	results := fanIn(workerResult...)

	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("done")

}
