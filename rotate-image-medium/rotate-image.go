package main

import "fmt"

func rotate(matrix [][]int) {
	fmt.Println(matrix)
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[len(matrix)-i-1] = matrix[len(matrix)-i-1], matrix[i]
	}
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
func main() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
}
