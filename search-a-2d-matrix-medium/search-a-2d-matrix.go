package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)*len(matrix[0])-1
	for left <= right {
		middle := (right + left) / 2
		i, j := middle/len(matrix[0]), middle%len(matrix[0])
		//fmt.Printf("mid:%d, row:%d, col:%d \n", middle, i, j)
		if matrix[i][j] == target {
			//fmt.Println(i, j, matrix[i][j])
			return true
		} else if matrix[i][j] < target {
			left = middle + 1
		} else if matrix[i][j] > target {
			right = middle - 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60}}
	target := 1
	fmt.Println(searchMatrix(matrix, target))
}
