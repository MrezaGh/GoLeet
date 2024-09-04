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

func swimInWater(grid [][]int) int {
	graph := make(map[int][]int)
	for i, row := range grid {
		for j, col := range row {
			if i < len(grid[0])-1 {
				graph[col] = append(graph[col], grid[i+1][j])
			}
			if i > 0 {
				graph[col] = append(graph[col], grid[i-1][j])
			}
			if j < len(grid)-1 {
				graph[col] = append(graph[col], grid[i][j+1])
			}
			if j > 0 {
				graph[col] = append(graph[col], grid[i][j-1])
			}
		}
	}
	//fmt.Println(graph)
	//build graph
	// bfs
	// minHeap frontier
	// in bfs:
	//check if target return
	// else open next <time> (a.k.a continue)
	frontier := MinHeap{}
	heap.Push(&frontier, grid[0][0])
	target := grid[len(grid)-1][len(grid)-1]
	visited := make(map[int]bool)
	current_level := grid[0][0]
	for len(frontier) > 0 {
		curr := heap.Pop(&frontier).(int)
		if _, ok := visited[curr]; ok {
			continue
		}
		//fmt.Println(curr, current_level)
		if curr > current_level {
			current_level = curr
		}
		visited[curr] = true
		if curr == target {
			return current_level
		}
		//fmt.Println("neighbors:", graph[curr])
		for _, neighbor := range graph[curr] {
			if _, ok := visited[neighbor]; !ok {
				heap.Push(&frontier, neighbor)
			}
		}
	}

	return current_level
}
func main() {

	grid := [][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}
	fmt.Println(swimInWater(grid))
}
