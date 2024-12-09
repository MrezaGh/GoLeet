package primitives

import (
	"fmt"
	"sync"
	"time"
)

func queuesim() {
	cond := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		cond.L.Lock()
		queue = queue[1:]
		fmt.Println("removing from queue")
		cond.L.Unlock()
		cond.Signal()

	}

	for i := 0; i < 10; i++ {
		cond.L.Lock()
		for len(queue) == 2 {
			cond.Wait()
		}
		queue = append(queue, i)
		fmt.Println("added to queue")
		go removeFromQueue(1 * time.Second)
		cond.L.Unlock()
	}
}

type Button struct {
	Clicked *sync.Cond
}

func btnSim() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}
	subscribe := func(cond *sync.Cond, fn func()) {
		var runningGoRoutines sync.WaitGroup
		runningGoRoutines.Add(1)
		go func() {
			runningGoRoutines.Done()
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			fn()
		}()
		runningGoRoutines.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("action 1")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("action 2")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("action 3")
		clickRegistered.Done()
	})
	//time.Sleep(1 * time.Nanosecond)
	button.Clicked.Broadcast()

	clickRegistered.Wait()

}

func main() {
	//queuesim()
	btnSim()
}
