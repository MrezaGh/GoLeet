package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Item struct {
	val  int
	i, j int
}

func longestIncreasingPath(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	items := make([]Item, 0, len(matrix)*len(matrix[0]))
	for i, row := range matrix {
		for j, col := range row {
			items = append(items, Item{
				val: col,
				i:   i,
				j:   j,
			})
		}
	}

	slices.SortFunc(items, func(a, b Item) int {
		return cmp.Compare(a.val, b.val)
	})
	slices.Reverse(items)
	for _, item := range items {
		fillNeighbours(dp, item, matrix)
	}

	totalMax := 0
	for _, row := range dp {
		totalMax = max(totalMax, slices.Max(row))
	}
	return totalMax + 1 //need <+1> because dp starts from 0 not 1
}

func fillNeighbours(dp [][]int, item Item, matrix [][]int) {
	directions := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for _, direction := range directions {
		newI, newJ := item.i+direction[0], item.j+direction[1]

		if newI < 0 || newI >= len(dp) || newJ < 0 || newJ >= len(dp[0]) {
			continue
		}

		if item.val > matrix[newI][newJ] {
			dp[newI][newJ] = max(dp[newI][newJ], dp[item.i][item.j]+1)
		}
	}

}

func main() {
	//matrix := [][]int{
	//	{9, 9, 4},
	//	{6, 6, 8},
	//	{2, 1, 1},
	//}

	matrix := [][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}

	fmt.Println(longestIncreasingPath(matrix))
}
