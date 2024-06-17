package main

import (
	"fmt"
)

func main() {
	n := 6
	edges := [][]int{
		{0, 1},
		{0, 2},
		{3, 5},
		{5, 4},
		{4, 3},
	}
	source := 0
	destination := 5
	fmt.Println(validPath(n, edges, source, destination))
	//queue := Queue{}
	//queue.Initialize(n)
	//fmt.Println(queue.IsEmpty())
	//queue.Enqueue(1)
	//fmt.Println(queue.IsEmpty())
	//queue.Enqueue(2)
	//queue.Enqueue(3)
	//fmt.Println(queue.Dequeue())
	//fmt.Println(queue.Dequeue())
	//queue.Enqueue(4)
	//fmt.Println(queue.Dequeue())
	//

}

type Queue struct {
	queue    []int
	head     int
	tail     int
	length   int
	capacity int
}

func (q *Queue) Initialize(n int) {
	q.head = 0
	q.tail = -1
	q.length = 0
	q.queue = make([]int, n+1)
	q.capacity = n
}

func (q *Queue) IsEmpty() bool {
	//return q.tail < q.head
	return q.length <= 0
}

func (q *Queue) Enqueue(i int) {
	q.tail += 1
	q.tail %= q.capacity
	q.queue[q.tail] = i
	q.length += 1

}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	} else {
		q.length -= 1
		element := q.queue[q.head]
		q.head += 1
		q.head %= q.capacity
		return element, true
	}
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	//fmt.Println("graph", graph)
	frontier := Queue{}
	frontier.Initialize(n)
	seen := make(map[int]struct{})

	frontier.Enqueue(source)
	seen[source] = struct{}{}

	for !frontier.IsEmpty() {
		node, _ := frontier.Dequeue()
		//fmt.Println("first", node, frontier, seen)
		seen[node] = struct{}{}

		if node == destination {
			return true
		}

		nextNodes := graph[node]
		for _, nextNode := range nextNodes {
			if _, exists := seen[nextNode]; exists {
				continue
			}
			frontier.Enqueue(nextNode)

		}

		//fmt.Println("end", node, frontier, seen)

	}

	//fmt.Println(frontier, graph)
	return false
}
