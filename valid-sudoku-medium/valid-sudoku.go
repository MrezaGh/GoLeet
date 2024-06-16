package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		rowMap := make(map[byte]struct{})
		colMap := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			// checking row
			if board[i][j] != '.' {
				if _, ok := rowMap[board[i][j]]; ok {
					return false
				} else {
					rowMap[board[i][j]] = struct{}{}
				}
			}

			//checking col
			if board[j][i] != '.' {
				if _, ok := colMap[board[j][i]]; ok {
					return false
				} else {
					colMap[board[j][i]] = struct{}{}
				}
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			box := make(map[byte]struct{})
			for i2 := 0; i2 < 3; i2++ {
				for j2 := 0; j2 < 3; j2++ {
					if board[i+i2][j+j2] == '.' {
						continue
					}
					if _, ok := box[board[i+i2][j+j2]]; ok {
						//fmt.Println("!!!!", string(board[i+i2][j+j2]), i+i2, j+j2)
						return false
					} else {
						box[board[i+i2][j+j2]] = struct{}{}
					}
				}
			}
		}
	}

	return true
}

func main() {
	board := [][]byte{
		//{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'9', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(board))
}
