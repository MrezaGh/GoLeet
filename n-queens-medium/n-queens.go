package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {

	ans := make([][]string, 0)

	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	var backtrack func(row int) bool

	backtrack = func(row int) bool {
		if row == n {
			validBoard := make([]string, len(board))
			for i := 0; i < len(board); i++ {
				validBoard[i] = strings.Join(board[i], "")
			}
			ans = append(ans, validBoard)

			return true
		}

		for j := 0; j < n; j++ {
			board[row][j] = "Q"
			if !checkBoard(board) || !backtrack(row+1) {
				board[row][j] = "."
				continue
			}
			board[row][j] = "."
		}

		return false
	}

	backtrack(0)

	return ans
}

func checkBoard(board [][]string) bool {
	rowCount := make([]int, len(board))
	columnCount := make([]int, len(board))
	diaCount := make([]int, 2*len(board))
	revDiaCount := make([]int, 2*len(board))

	for i, _ := range board {
		for j := 0; j < len(board); j++ {
			if board[i][j] == "Q" {
				if rowCount[i] >= 1 || columnCount[j] >= 1 || diaCount[(i-j)+len(board)] >= 1 || revDiaCount[i+j] >= 1 {
					return false
				}
				rowCount[i] += 1
				columnCount[j] += 1
				diaCount[(i-j)+len(board)] += 1
				revDiaCount[i+j] += 1
			}
		}
	}
	return true
}

func main() {
	n := 6
	//n := 1

	fmt.Println(solveNQueens(n), len(solveNQueens(n)))
}
