package main

import "fmt"

func solve(board [][]byte) {
	frontier := make([][2]int, 0)
	edgeAccess := make(map[[2]int]bool)

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, i := range [2]int{0, len(board) - 1} {
		for j := 0; j < len(board[0]); j++ {
			if !edgeAccess[[2]int{i, j}] && board[i][j] == 'O' {
				frontier = append(frontier, [2]int{i, j})
				edgeAccess[[2]int{i, j}] = true
			}
		}
	}
	for _, j := range [2]int{0, len(board[0]) - 1} {
		for i := 0; i < len(board); i++ {
			if !edgeAccess[[2]int{i, j}] && board[i][j] == 'O' {
				frontier = append(frontier, [2]int{i, j})
				edgeAccess[[2]int{i, j}] = true
			}
		}
	}

	for len(frontier) > 0 {
		head := frontier[0]
		frontier = frontier[1:]
		for _, dir := range dirs {
			newI, newJ := head[0]+dir[0], head[1]+dir[1]
			if newI >= 0 && newI < len(board) && newJ >= 0 && newJ < len(board[0]) && board[newI][newJ] == 'O' && !edgeAccess[[2]int{newI, newJ}] {
				edgeAccess[[2]int{newI, newJ}] = true
				frontier = append(frontier, [2]int{newI, newJ})
			}
		}
	}

	for i, row := range board {
		for j, col := range row {
			if col == 'O' && !edgeAccess[[2]int{i, j}] {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	fmt.Println(board)
}
