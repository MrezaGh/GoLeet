package pattern

import (
	"fmt"
	"math/rand"
	"time"
)

func leak() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("did the work")
			defer close(completed)
			for s := range strings {
				fmt.Println("received:", s)
			}
		}()
		return completed
	}

	doWork(nil)
	fmt.Println("all done")
}

func noLeakWithCancel() {
	doWork := func(done <-chan int, strings <-chan string) <-chan int {
		terminated := make(chan int)
		go func() {
			defer fmt.Println("exiting dowork")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println("found:", s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("triggering done")
		close(done)
	}()

	<-doWork(done, nil)
	fmt.Println("all done")
}

func randStream() {
	newRandStream := func(done <-chan int) <-chan int {
		randStreamChan := make(chan int)
		go func() {
			defer println("exiting random generator")
			defer close(randStreamChan)
			for {
				select {
				case randStreamChan <- rand.Int():
				case <-done:
					return
				}
			}
		}()

		return randStreamChan
	}
	done := make(chan int)
	rands := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 0; i < 3; i++ {
		fmt.Println(<-rands)
	}
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("all done")

}
