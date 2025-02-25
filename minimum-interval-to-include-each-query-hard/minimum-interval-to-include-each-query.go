package main

import (
	"cmp"
	"container/heap"
	"fmt"
	"slices"
)

type MinHeap [][2]int // [length, lastIndex]

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][1]-h[i][0] < h[j][1]-h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minInterval(intervals [][]int, queries []int) []int {
	indexedQueries := make([][2]int, len(queries))
	for i := 0; i < len(queries); i++ {
		indexedQueries[i] = [2]int{queries[i], i}
	}
	slices.SortFunc(indexedQueries, func(a, b [2]int) int {
		return cmp.Compare(a[0], b[0])
	})
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	slices.Sort(queries)
	minHeap := MinHeap{}
	heap.Init(&minHeap)
	ans := make([]int, 0, len(queries))

	j := 0
	for _, interval := range intervals {
		for j < len(queries) && interval[0] > queries[j] {
			for minHeap.Len() > 0 && minHeap[0][1] < queries[j] {
				heap.Pop(&minHeap)
			}
			if minHeap.Len() > 0 {
				head := minHeap[0]
				ans = append(ans, head[1]-head[0]+1)
			} else {
				ans = append(ans, -1)
			}
			j++
		}
		heap.Push(&minHeap, [2]int{interval[0], interval[1]})
	}

	for ; j < len(queries); j++ {
		for minHeap.Len() > 0 && minHeap[0][1] < queries[j] {
			heap.Pop(&minHeap)
		}
		if minHeap.Len() > 0 {
			head := minHeap[0]
			ans = append(ans, head[1]-head[0]+1)
		} else {
			ans = append(ans, -1)
		}
	}

	originalAns := make([]int, len(queries))
	for i, query := range indexedQueries {
		originalAns[query[1]] = ans[i]
	}
	return originalAns
}

func main() {
	//intervals, queries := [][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5}
	intervals, queries := [][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22}
	fmt.Println(minInterval(intervals, queries))
}
