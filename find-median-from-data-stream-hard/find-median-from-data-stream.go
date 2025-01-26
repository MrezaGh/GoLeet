package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	minHeap MinHeap
	maxHeap MinHeap
}

func Constructor() MedianFinder {
	mf := MedianFinder{
		minHeap: MinHeap{},
		maxHeap: MinHeap{},
	}
	heap.Init(&mf.minHeap)
	heap.Init(&mf.maxHeap)
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	// FIXME: NEGATE for maxHeap
	//biggerSide := this.minHeap[0]
	if this.minHeap.Len() == 0 && this.maxHeap.Len() == 0 {
		heap.Push(&this.maxHeap, -num)
		return
	}

	lesserSide := -this.maxHeap[0]
	if num > lesserSide {
		heap.Push(&this.minHeap, num)
	} else {
		heap.Push(&this.maxHeap, -num)
	}

	if this.minHeap.Len() > this.maxHeap.Len() {

		heap.Push(&this.maxHeap, -heap.Pop(&this.minHeap).(int))
	} else if this.maxHeap.Len()-this.minHeap.Len() > 1 {
		heap.Push(&this.minHeap, -heap.Pop(&this.maxHeap).(int))
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64(-this.maxHeap[0])
	} else if this.maxHeap.Len() < this.minHeap.Len() {
		return float64(this.minHeap[0])
	} else {
		return float64(this.minHeap[0]-this.maxHeap[0]) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	medianFinder := Constructor()

	medianFinder.AddNum(1)                 // arr = [1]
	medianFinder.AddNum(2)                 // arr = [1, 2]
	fmt.Println(medianFinder.FindMedian()) // return 1.5 (i.e., (1 + 2) / 2)
	medianFinder.AddNum(3)                 // arr[1, 2, 3]
	fmt.Println(medianFinder.FindMedian()) // return 2.0
}
