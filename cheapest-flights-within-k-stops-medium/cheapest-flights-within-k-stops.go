package main

import (
	"fmt"
	"math"
)

//type MinHeap [][4]int //[Node, price, stops, prev]
//
//func (h MinHeap) Len() int           { return len(h) }
//func (h MinHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
//func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
//func (h *MinHeap) Push(x any) {
//	// Push and Pop use pointer receivers because they modify the slice's length,
//	// not just its contents.
//	*h = append(*h, x.([4]int))
//}
//
//func (h *MinHeap) Pop() any {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
//
//func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
//	graph := make(map[int][][2]int) // [node, weight]
//	for _, flight := range flights {
//		graph[flight[0]] = append(graph[flight[0]], [2]int{flight[1], flight[2]})
//	}
//
//	visited := make(map[[3]int]bool)
//	frontier := MinHeap{}
//	heap.Push(&frontier, [4]int{src, 0, -1, src})
//
//	for len(frontier) > 0 {
//		curr := heap.Pop(&frontier).([4]int)
//		node, curr_price, stops, prev := curr[0], curr[1], curr[2], curr[3]
//		if _, ok := visited[[3]int{node, prev, stops}]; ok {
//			//fmt.Println("skipping:", node, prev)
//			continue
//		}
//		visited[[3]int{node, prev, stops}] = true
//		//fmt.Println("current node:", node, curr_price, stops)
//		if node == dst && stops <= k {
//			return curr_price
//		}
//		if stops >= k {
//			continue
//		}
//		for _, neighbor := range graph[node] {
//			heap.Push(&frontier, [4]int{neighbor[0], curr_price + neighbor[1], stops + 1, node})
//		}
//
//	}
//
//	return -1
//}

// faster solution with DP
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt
	}

	cost[src] = 0
	for i := 0; i <= k; i++ {
		temp := make([]int, n)
		copy(temp, cost)
		for _, flight := range flights {
			from, to, price := flight[0], flight[1], flight[2]
			if cost[from] != math.MaxInt {
				//fmt.Println("here->", from, to, price, cost[from])
				//fmt.Println(cost[to], cost[from]+price)
				temp[to] = min(temp[to], cost[from]+price)
			}
		}
		cost = temp
	}

	if cost[dst] == math.MaxInt {
		return -1
	}
	return cost[dst]
}

func main() {
	//n, src, dst, k := 4, 0, 3, 1
	//flights := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}

	n, src, dst, k := 3, 0, 2, 1
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}

	//n, src, dst, k := 3, 0, 2, 0
	//flights := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}

	//n, src, dst, k := 4, 0, 3, 1
	//flights := [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}

	//n, src, dst, k := 5, 0, 2, 2
	//flights := [][]int{{0, 1, 5}, {1, 2, 5}, {0, 3, 2}, {3, 1, 2}, {1, 4, 1}, {4, 2, 1}}
	fmt.Println(findCheapestPrice(n, flights, src, dst, k))
}
