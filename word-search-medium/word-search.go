package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visited := make(map[[2]int]bool)

	var backtrack func(row, column int, target string) bool
	backtrack = func(i, j int, target string) bool {
		if target == "" {
			return true
		}
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return false
		}
		//fmt.Printf("checking from: i,j=%d,%d  , target=%s\n", i, j, word)

		if board[i][j] != target[0] || visited[[2]int{i, j}] {
			//fmt.Println("cont")
			return false
		}
		visited[[2]int{i, j}] = true
		for _, dir := range directions {
			if backtrack(i+dir[0], j+dir[1], target[1:]) {
				return true
			}
		}
		visited[[2]int{i, j}] = false

		return false
	}
	for i, _ := range board {
		for j := 0; j < len(board[0]); j++ {
			if backtrack(i, j, word) {
				return true
			}
		}
	}
	return false
}

func main() {
	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//word := "ABCCED"

	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//word := "SEE"
	//
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCB"

	fmt.Println(exist(board, word))
}
