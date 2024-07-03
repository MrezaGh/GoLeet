package main

import (
	"fmt"
)

var directions = [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func pacificAtlantic(heights [][]int) [][]int {
	var frontier [][2]int
	pacific := make(map[[2]int]bool)
	atlantic := make(map[[2]int]bool)
	for i := 0; i < len(heights); i++ {
		frontier = append(frontier, [2]int{i, 0})
		pacific[[2]int{i, 0}] = true

	}
	for j := 0; j < len(heights[0]); j++ {
		frontier = append(frontier, [2]int{0, j})
		pacific[[2]int{0, j}] = true
	}
	bfs(heights, frontier, pacific)
	frontier = [][2]int{}
	for i := 0; i < len(heights); i++ {
		frontier = append(frontier, [2]int{i, len(heights[0]) - 1})
		atlantic[[2]int{i, len(heights[0]) - 1}] = true

	}
	for j := 0; j < len(heights[0]); j++ {
		frontier = append(frontier, [2]int{len(heights) - 1, j})
		atlantic[[2]int{len(heights) - 1, j}] = true
	}
	bfs(heights, frontier, atlantic)
	//fmt.Println(pacific)
	//fmt.Println(atlantic)
	var result [][]int
	for key, _ := range pacific {
		if _, ok := atlantic[key]; ok {
			result = append(result, []int{key[0], key[1]})
		}
	}

	return result
}

func bfs(heights [][]int, frontier [][2]int, ocean map[[2]int]bool) {
	for len(frontier) > 0 {
		head := frontier[0]
		frontier = frontier[1:]
		for _, dir := range directions {
			i, j := dir[0]+head[0], dir[1]+head[1]
			if _, ok := ocean[[2]int{i, j}]; !ok && i >= 0 && i < len(heights) && j >= 0 && j < len(heights[0]) &&
				heights[i][j] >= heights[head[0]][head[1]] {
				ocean[[2]int{i, j}] = true
				frontier = append(frontier, [2]int{i, j})
			}
		}
	}
}

func main() {
	heights := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	//heights := [][]int{{1}}
	fmt.Println(pacificAtlantic(heights))
}
