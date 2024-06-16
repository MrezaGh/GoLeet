package main

import (
	"container/heap"
	"fmt"
	"math"
)

type MaxHeap [][2]int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] > h[j][0]
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([2]int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func maxSlidingWindow(nums []int, k int) []int {
	maxWindow := make([]int, 0)
	h := &MaxHeap{}
	left := 0
	currentMax := nums[0]
	head := [2]int{0, -1}
	//var head [2]int
	for right := 0; right < len(nums); right++ {
		num := nums[right]
		heap.Push(h, [2]int{num, right})
		currentMax = max(currentMax, num)

		if head == [2]int{0, -1} {
			head = heap.Pop(h).([2]int)
			currentMax = head[0]
			//fmt.Println("first left move:", head)
		}
		if right >= k-1 {
			maxWindow = append(maxWindow, currentMax)
		}
		if right-left+1 == k {
			leftItem := [2]int{nums[left], left}
			if leftItem == head {
				for h.Len() > 0 {
					head = heap.Pop(h).([2]int)
					if head[1] >= left {
						currentMax = head[0]
						//fmt.Println("moving left:", head)
						break
					}
				}
				if h.Len() == 0 {
					currentMax = math.MinInt
				}
			}

			//fmt.Println(head)
			left += 1
		}
	}

	return maxWindow
}

func main() {
	//nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	//k := 3
	nums := []int{1}
	k := 1
	//nums := []int{1, -1}
	//k := 1

	fmt.Println(maxSlidingWindow(nums, k))
}
