package main

import (
	"fmt"
	"slices"
)

func mergeTriplets(triplets [][]int, target []int) bool {
	merged := []int{0, 0, 0}

	for _, triplet := range triplets {
		if triplet[0] > target[0] || triplet[1] > target[1] || triplet[2] > target[2] {
			continue
		}
		merged[0] = max(merged[0], triplet[0])
		merged[1] = max(merged[1], triplet[1])
		merged[2] = max(merged[2], triplet[2])
	}

	if slices.Equal(merged, target) {
		return true
	}

	return false
}

func main() {
	//triplets, target := [][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5}
	triplets, target := [][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5}
	//triplets, target := [][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}}, []int{5, 5, 5}
	fmt.Println(mergeTriplets(triplets, target))
}
