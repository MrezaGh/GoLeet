package main

import "fmt"

//func findRedundantConnection(edges [][]int) []int {
//	graph := make(map[int][]int)
//	degree := make(map[int]int)
//	for _, edge := range edges {
//		node1, node2 := edge[0], edge[1]
//		graph[node1] = append(graph[node1], node2)
//		graph[node2] = append(graph[node2], node1)
//		degree[node1] += 1
//		degree[node2] += 1
//	}
//	for i := len(edges) - 1; i >= 0; i-- {
//		if degree[edges[i][0]]-1 == 0 || degree[edges[i][1]]-1 == 0 {
//			continue
//		}
//		fmt.Println("before remove:", "i,j:", edges[i], graph[edges[i][0]], graph[edges[i][1]])
//		for j, edge := range graph[edges[i][0]] {
//			if edge == edges[i][1] {
//				graph[edges[i][0]] = append(graph[edges[i][0]][:j], graph[edges[i][0]][j+1:]...)
//			}
//		}
//		for j, edge := range graph[edges[i][1]] {
//			if edge == edges[i][0] {
//				graph[edges[i][1]] = append(graph[edges[i][1]][:j], graph[edges[i][1]][j+1:]...)
//			}
//		}
//		fmt.Println("after remove:", edges[i], graph[edges[i][0]], graph[edges[i][1]])
//		if !hasCycle(graph) {
//			return edges[i]
//		}
//		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
//		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
//	}
//	// remove edge
//	//see if has cycle
//	// add edge
//	return nil
//}
//
//func hasCycle(graph map[int][]int) bool {
//	visited := make(map[int]bool)
//	degrees := make(map[int]int)
//	queue := make([]int, 0)
//	for node, neighbors := range graph {
//		degrees[node] = len(neighbors)
//		if len(neighbors) == 1 {
//			queue = append(queue, node)
//		}
//	}
//
//	for len(queue) > 0 {
//		head := queue[0]
//		if _, exist := visited[head]; exist {
//			queue = queue[1:]
//			continue
//		}
//		visited[head] = true
//		for _, neighbor := range graph[head] {
//			if _, exist := visited[neighbor]; !exist {
//				degrees[neighbor] -= 1
//				degrees[head] -= 1
//				if degrees[neighbor] == 1 {
//					queue = append(queue, neighbor)
//				}
//			}
//		}
//		queue = queue[1:]
//	}
//	fmt.Println("cycle:", degrees)
//	for _, degree := range degrees {
//		if degree != 0 {
//			return true
//		}
//	}
//	return false
//}

// faster approach with Disjoint set union =============
type DSU struct {
	parent [1001]int
}

func (d *DSU) find(node int) int {
	if d.parent[node] != node {
		d.parent[node] = d.find(d.parent[node])
	}
	return d.parent[node]
}

func (d *DSU) union(node1, node2 int) bool {
	if d.find(node1) == d.find(node2) {
		return false
	} else {
		d.parent[node1] = d.find(node2)
	}
	return true
}

func (d *DSU) add(node int) {
	if d.parent[node] == 0 {
		d.parent[node] = node
	}
}

func findRedundantConnection(edges [][]int) []int {
	dsu := &DSU{}

	for _, edge := range edges {
		dsu.add(edge[0])
		dsu.add(edge[1])
		if !dsu.union(dsu.find(edge[0]), dsu.find(edge[1])) {
			return edge
		}
	}
	return nil
}

func main() {
	//edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
	//edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	//edges := [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}, {1, 5}}
	edges := [][]int{{9, 10}, {5, 8}, {2, 6}, {1, 5}, {3, 8}, {4, 9}, {8, 10}, {4, 10}, {6, 8}, {7, 9}}
	fmt.Println(findRedundantConnection(edges))
}
