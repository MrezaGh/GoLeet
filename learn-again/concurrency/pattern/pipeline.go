package pattern

import (
	"fmt"
	"math/rand"
)

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	stream := make(chan interface{})
	go func() {
		defer close(stream)
		for {
			for _, val := range values {
				select {
				case <-done:
					return
				case stream <- val:
				}
			}
		}
	}()

	return stream
}

func repeatFN(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	stream := make(chan interface{})
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func take(done <-chan interface{}, stream <-chan interface{}, num int) <-chan interface{} {
	outputStream := make(chan interface{})
	go func() {
		defer close(outputStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case outputStream <- <-stream:
			}
		}
	}()
	return outputStream
}

func UseRepeatTakePipeline() {

	done := make(chan interface{})
	defer close(done)

	//data := []interface{}{1, 2, 3}
	//for i := range take(done, repeat(done, data...), 10) {
	//	fmt.Printf("%d ", i)
	//}

	fn := func() interface{} { return rand.Int() }
	for i := range take(done, repeatFN(done, fn), 10) {
		fmt.Printf("%d ", i)
	}

}
