package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	seen := make(map[[2]int]bool)
	cur := [2]int{0, 0}
	output := []int{matrix[0][0]}
	seen[cur] = true
	counter := 1
outer:
	for {
		for _, dir := range directions {
			next := [2]int{cur[0] + dir[0], cur[1] + dir[1]}
			for next[0] >= 0 && next[0] < len(matrix) && next[1] >= 0 && next[1] < len(matrix[0]) {
				if _, exist := seen[next]; exist {
					break
				}
				counter += 1
				cur = next
				//fmt.Println(cur)
				output = append(output, matrix[cur[0]][cur[1]])
				seen[cur] = true

				next = [2]int{cur[0] + dir[0], cur[1] + dir[1]}
			}
			if counter >= len(matrix)*len(matrix[0]) {
				//fmt.Println("done")
				break outer
			}
		}
	}

	return output
}

func main() {

	//matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
}
