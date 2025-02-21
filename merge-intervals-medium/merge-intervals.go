package main

import (
	"cmp"
	"fmt"
	"slices"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	merged := make([][]int, 0)
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		head := merged[len(merged)-1]
		if head[1] >= intervals[i][0] {
			merged[len(merged)-1] = []int{head[0], max(intervals[i][1], head[1])}
		} else {
			merged = append(merged, intervals[i])
		}
	}

	return merged
}
func main() {
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals := [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals))
}
