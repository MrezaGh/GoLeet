package main

import "fmt"

func numIslands(grid [][]byte) int {
	visited := make(map[[2]int]bool)
	graph := make(map[[2]int][][2]int)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if string(grid[i][j]) == "0" {
				continue
			}
			key := [2]int{i, j}
			if i < len(grid)-1 && grid[i+1][j] == '1' {
				graph[key] = append(graph[key], [2]int{i + 1, j})
			}
			if j < len(grid[0])-1 && grid[i][j+1] == '1' {
				graph[key] = append(graph[key], [2]int{i, j + 1})
			}
			if i > 0 && grid[i-1][j] == '1' {
				graph[key] = append(graph[key], [2]int{i - 1, j})
			}
			if j > 0 && grid[i][j-1] == '1' {
				graph[key] = append(graph[key], [2]int{i, j - 1})
			}
		}
	}

	//fmt.Println(visited, graph)
	islands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '0' {
				continue
			}
			key := [2]int{i, j}
			if _, ok := visited[key]; ok {
				continue
			} else {
				fmt.Println("found new island:", key)
				islands += 1
				visited[key] = true
				bfs(graph, key, visited)
			}
		}
	}

	return islands
}

func bfs(graph map[[2]int][][2]int, node [2]int, visited map[[2]int]bool) {
	for _, n := range graph[node] {
		if _, ok := visited[n]; ok {
			continue
		}
		visited[n] = true
		bfs(graph, n, visited)
	}
}

func main() {
	//grid := [][]byte{
	//	{'1', '1', '1', '1', '0'},
	//	{'1', '1', '0', '1', '0'},
	//	{'1', '1', '0', '0', '0'},
	//	{'0', '0', '0', '0', '0'},
	//}

	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}
