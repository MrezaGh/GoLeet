package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	maxHeap := MaxHeap{}
	heap.Init(&maxHeap)
	for _, stone := range stones {
		heap.Push(&maxHeap, stone)
	}
	for maxHeap.Len() > 0 {
		if maxHeap.Len() == 1 {
			return heap.Pop(&maxHeap).(int)
		}
		first, second := heap.Pop(&maxHeap), heap.Pop(&maxHeap)
		if first == second {
			continue
		} else {
			newStone := first.(int) - second.(int)
			heap.Push(&maxHeap, newStone)
		}
	}
	return 0
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones))
}
