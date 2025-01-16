package main

import (
	"cmp"
	"fmt"
	"slices"
)

func kClosest(points [][]int, k int) [][]int {
	distances := make([][2]int, 0) // [distances to origin, point index]

	for i, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		distances = append(distances, [2]int{distance, i})
	}

	slices.SortedFunc(slices.Values(distances), func(e [2]int, e2 [2]int) int {
		return cmp.Compare(e[0], e2[0])
	})
	fmt.Println(distances)

	return nil
}

func main() {
	points, k := [][]int{{1, 3}, {-2, 2}}, 1
	//points, k := [][]int{{3,3},{5,-1},{-2,4}}, 2

	fmt.Println(kClosest(points, k))
}
