package main

import (
	"cmp"
	"fmt"
	"slices"
)

func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	erased := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < erased[len(erased)-1][1] {
			if intervals[i][1] < erased[len(erased)-1][1] {
				erased[len(erased)-1] = intervals[i]
			}
		} else {
			erased = append(erased, intervals[i])
		}
	}
	//fmt.Println(erased)
	return len(intervals) - len(erased)
}
func main() {
	//intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	intervals := [][]int{{1, 2}, {1, 2}, {1, 2}}
	fmt.Println(eraseOverlapIntervals(intervals))
}
