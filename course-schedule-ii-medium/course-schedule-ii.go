package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	adjacency := make(map[int][]int)
	degrees := make(map[int]int)
	for i := 0; i < numCourses; i++ {
		degrees[i] = 0
		adjacency[i] = make([]int, 0)
	}

	for _, prereq := range prerequisites {
		course, req := prereq[0], prereq[1]
		degrees[course] += 1
		adjacency[req] = append(adjacency[req], course)
	}
	//fmt.Println(degrees)
	//fmt.Println(adjacency)
	bfs := make([]int, 0)
	for node, degree := range degrees {
		if degree == 0 {
			bfs = append(bfs, node)
		}
	}
	if len(bfs) == 0 {
		return []int{}
	}
	//fmt.Println("init bfs", bfs)
	index := 0
	for index < len(bfs) {
		node := bfs[index]
		//fmt.Println("current node:", node)
		for _, neighbour := range adjacency[node] {
			degrees[neighbour] -= 1
			if degrees[neighbour] == 0 {
				bfs = append(bfs, neighbour)
			}
		}
		//fmt.Println("bfs:", bfs)
		index += 1
	}
	//fmt.Println(degrees)
	if len(bfs) == numCourses {
		return bfs
	}
	return []int{}
}

func main() {
	//numCourses, prerequisites := 2, [][]int{{1, 0}}
	//numCourses, prerequisites := 4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	numCourses, prerequisites := 5, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}, {4, 3}, {3, 4}}
	fmt.Println(findOrder(numCourses, prerequisites))
}
