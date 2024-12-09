package pattern

func orDone(done <-chan interface{}, ch <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		close(valStream)
		for {
			select {
			case <-done:
				return
			case i, ok := <-ch:
				if ok == false {
					return
				}
				select {
				case valStream <- i:
				case <-done:
					return
				}
			}
		}
	}()
	return valStream
}

func bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		for {
			stream := make(<-chan interface{})
			potentialStream, ok := <-chanStream
			if ok == false {
				return
			}
			stream = potentialStream
			for val := range orDone(done, stream) {
				select {
				case <-done:
				case valStream <- val:
				}
			}

		}

	}()
	return valStream
}

func UseOrDone() {

}
