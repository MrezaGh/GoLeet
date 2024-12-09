package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Priority int
	fn       func() interface{}
}

type Result struct {
	TaskId int
}

type Worker struct {
	ID    int
	tasks <-chan Task
}

func (w *Worker) Run(wg *sync.WaitGroup) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		for task := range w.tasks {
			fmt.Printf("worker:%d doing task:%d\n", w.ID, task.ID)
			out <- w.execute(task)
			wg.Done()
		}
	}()

	return out
}

func (w *Worker) execute(task Task) interface{} {
	return task.fn()
}

type Queue interface {
	Add(task Task, priority int) int
	Cancel(taskID int) error
	GetStatus(taskID int) interface{}
	GetResult(taskID int) interface{}
}

type ConcurrentTaskQueue struct {
	wg            sync.WaitGroup
	tasks         chan Task
	workers       []Worker
	workerResults []<-chan interface{}
	idCounter     int
}

func NewConcurrentTaskQueue(numWorkers int) *ConcurrentTaskQueue {
	q := &ConcurrentTaskQueue{
		tasks:         make(chan Task),
		workers:       make([]Worker, numWorkers),
		workerResults: make([]<-chan interface{}, numWorkers),
		idCounter:     1,
	}

	for i := 0; i < numWorkers; i++ {
		q.workers[i] = Worker{
			ID:    i,
			tasks: q.tasks,
		}
		q.workerResults[i] = q.workers[i].Run(&q.wg)
	}

	return q

}

func (q *ConcurrentTaskQueue) Add(fn func() interface{}, priority int) int {
	q.wg.Add(1)
	task := Task{
		ID:       q.idCounter,
		Priority: priority,
		fn:       fn,
	}
	q.idCounter++ // TODO: Should be trade safe
	go func() {
		q.tasks <- task

	}()

	return task.ID
}

func (q *ConcurrentTaskQueue) fanIn() <-chan interface{} {
	results := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(q.workerResults))

	go func() {
		defer close(results)
		for _, resultChan := range q.workerResults {

			go func(resultChan <-chan interface{}) {
				defer wg.Done()
				for item := range resultChan {
					results <- item
				}
			}(resultChan)

		}
		wg.Wait()
	}()

	return results
}

func main() {
	//var q Queue
	numWorkers := 3
	q := NewConcurrentTaskQueue(numWorkers)

	//jobs := make([]func() interface{}, 101)
	for i := 1; i <= 10; i++ {
		//bind := i
		w8 := func() interface{} {
			//jobs[i] = func() interface{} {
			fmt.Println("waiting milis:", i)
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Println("waiting milis:", i, " done.")
			return i
		}
		q.Add(w8, 0)
	}
	go func() {
		q.wg.Wait()
		close(q.tasks)
	}()
	//jobs[10]()
	//fmt.Println("wwww", q)

	for result := range q.fanIn() {
		fmt.Println(result)
	}

	fmt.Printf("done")

}
