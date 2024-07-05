package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	for _, preReq := range prerequisites {
		course, req := preReq[0], preReq[1]
		graph[course] = append(graph[course], req)
	}
	//fmt.Println(graph)
	visited := make(map[int]bool)
	for i := 0; i < numCourses; i++ {
		if _, ok := graph[i]; ok && !visited[i] {
			recStack := make(map[int]bool)
			if hasCycle := dfs(graph, i, visited, recStack); hasCycle {
				fmt.Println("found cycle on:", i, "path:", recStack)
				return false
			}
		}
	}
	return true
}

// determines if we have a cycle or not
func dfs(graph map[int][]int, currentNode int, visited map[int]bool, recStack map[int]bool) bool {

	visited[currentNode] = true
	recStack[currentNode] = true
	//fmt.Println("inside dfs", visited, recStack)
	for _, neighbor := range graph[currentNode] {
		if recStack[neighbor] {
			return true
		}
		if visited[neighbor] {
			continue
		}
		if dfs(graph, neighbor, visited, recStack) {
			return true
		}
		recStack[neighbor] = false
	}

	return false
}

func main() {
	//numCourses := 2
	//prerequisites := [][]int{{1, 0}}
	//numCourses := 2
	//prerequisites := [][]int{{1, 0}, {0, 1}}
	numCourses := 6
	//prerequisites := [][]int{{1, 0}, {2, 1}, {2, 3}, {3, 4}, {3, 5}} // True
	prerequisites := [][]int{{1, 0}, {2, 1}, {2, 3}, {3, 4}, {3, 5}, {4, 2}} // False
	fmt.Println(canFinish(numCourses, prerequisites))

}
