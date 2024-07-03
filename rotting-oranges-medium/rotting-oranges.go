package main

import "fmt"

func orangesRotting(grid [][]int) int {
	var frontier [][3]int // i, j , (currentTime)
	visited := make(map[[2]int]bool)
	adjacent := [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	for i, row := range grid {
		for j, col := range row {
			if col == 2 {
				frontier = append(frontier, [3]int{i, j, 0})
				visited[[2]int{i, j}] = true // FIXME?
			}
		}
	}

	maxTime := 0
	for len(frontier) > 0 {
		head := frontier[0]
		frontier = frontier[1:]
		i, j, currentTime := head[0], head[1], head[2]
		maxTime = max(maxTime, currentTime)
		for _, adj := range adjacent {
			newI, newJ := i+adj[0], j+adj[1]
			if _, ok := visited[[2]int{newI, newJ}]; !ok && newI >= 0 && newI < len(grid) && newJ >= 0 && newJ < len(grid[0]) && grid[newI][newJ] == 1 {
				grid[newI][newJ] = 2
				visited[[2]int{newI, newJ}] = true
				frontier = append(frontier, [3]int{newI, newJ, currentTime + 1})
			}
		}

	}

	for _, row := range grid {
		for _, col := range row {
			if col == 1 {
				return -1
			}
		}
	}

	return maxTime
}

func main() {
	//grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	//grid := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	grid := [][]int{{1, 2}}
	//grid := [][]int{{0, 2}}
	fmt.Println(orangesRotting(grid))
}
