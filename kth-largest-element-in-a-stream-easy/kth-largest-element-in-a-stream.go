package main

import (
	"container/heap"
	"fmt"
	"slices"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	heap  *IntHeap
	limit int
}

func Constructor(k int, nums []int) KthLargest {
	slices.Sort(nums)
	slices.Reverse(nums)
	topK := IntHeap{}
	topK = append(topK, nums[:min(len(nums), k)]...)
	heap.Init(&topK)
	return KthLargest{heap: &topK,
		limit: k}
}

func (this *KthLargest) Add(val int) int {
	if this.heap.Len() < this.limit {
		heap.Push(this.heap, val)
		return (*this.heap)[0]
	}
	head := (*this.heap)[0]
	if head >= val {
		return head
	} else {
		heap.Pop(this.heap)
		heap.Push(this.heap, val)
		return (*this.heap)[0]
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
	k := 3
	nums := []int{4, 5, 8, 2}
	vals := []int{3, 5, 10, 9, 4}

	obj := Constructor(k, nums)
	for _, val := range vals {
		fmt.Println(obj.Add(val))
	}
}
