package main

import (
	"cmp"
	"fmt"
	"slices"
)

func kClosest(points [][]int, k int) [][]int {
	distances := make([][]int, 0) // [distances to origin, point index]

	for i, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		distances = append(distances, []int{distance, i})
	}

	slices.SortFunc(distances, func(e []int, e2 []int) int {
		return cmp.Compare(e[0], e2[0])
	})

	//fmt.Println(distances)
	ans := make([][]int, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, points[distances[i][1]])
	}

	return ans
}

func main() {
	//points, k := [][]int{{1, 3}, {-2, 2}}, 1
	points, k := [][]int{{3, 3}, {5, -1}, {-2, 4}}, 2

	fmt.Println(kClosest(points, k))
}
