package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []Task

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].remainingCount > h[j].remainingCount }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Task))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Task struct {
	remainingCount int
	character      string
}

func leastInterval(tasks []byte, n int) int {
	eligible := make(map[int][]Task)
	cycle := 0
	remainingTasks := len(tasks)
	maxHeap := MaxHeap{}
	heap.Init(&maxHeap)

	taskCount := make(map[string]int)
	for _, t := range tasks {
		taskCount[string(t)] += 1
	}
	for key, value := range taskCount {
		heap.Push(&maxHeap, Task{
			remainingCount: value,
			character:      key,
		})
	}

	for remainingTasks > 0 {
		for _, task := range eligible[cycle] {
			heap.Push(&maxHeap, task)
		}
		cycle += 1
		if maxHeap.Len() <= 0 {
			//fmt.Print("idle", "->")
			continue
		}
		task := heap.Pop(&maxHeap).(Task)
		//fmt.Print(task.character, "->")
		task.remainingCount -= 1
		if task.remainingCount > 0 {
			eligible[cycle+n] = append(eligible[cycle+n], task)
		}
		remainingTasks -= 1
	}

	return cycle
}

func main() {
	tasks, n := []byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2
	//tasks, n := []byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1
	//tasks, n := []byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3
	fmt.Println(leastInterval(tasks, n))
}
