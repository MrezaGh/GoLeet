package main

import (
	"container/heap"
	"fmt"
)

type MinHeap [][2]int

func (mh MinHeap) Len() int { return len(mh) }

func (mh MinHeap) Less(i, j int) bool {
	return mh[i][0] < mh[j][0]
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MinHeap) Push(x any) {
	item := x.([2]int)
	*mh = append(*mh, item)
}

func (mh *MinHeap) Pop() any {
	old := *mh
	n := len(old)
	item := old[n-1]
	*mh = old[0 : n-1]
	return item
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make(map[int][][2]int) // [node, time_distance]
	for _, edge := range times {
		graph[edge[0]] = append(graph[edge[0]], [2]int{edge[1], edge[2]})
	}
	//fmt.Println(graph)
	frontier := MinHeap{} // [distance, node]
	heap.Push(&frontier, [2]int{0, k})
	//bfs from node k
	// frontier is a min head
	last_time := 0
	visited := make(map[int]bool)
	for len(frontier) > 0 {
		head := heap.Pop(&frontier).([2]int)
		distance, node := head[0], head[1]
		//fmt.Println("currentNode:", node, "neeeded time:", distance)
		if _, ok := visited[node]; ok {
			continue
		}

		last_time = distance
		visited[node] = true
		//fmt.Println("neighbors:", graph[node])
		for _, neighbor := range graph[node] {
			currentNode, currentTime := neighbor[0], neighbor[1]
			if _, ok := visited[currentNode]; !ok {
				heap.Push(&frontier, [2]int{currentTime + distance, currentNode})
			}
		}
	}

	if len(visited) != n {
		return -1
	}

	return last_time
}

func main() {
	//times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	//times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}, {4, 5, 5}}
	//n, k := 5, 2
	//times := [][]int{{1, 2, 1}}
	//n, k := 2, 2
	times := [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 2}}
	n, k := 3, 1

	fmt.Println(networkDelayTime(times, n, k))
}
