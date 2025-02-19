package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	left := newInterval[0]
	leftIndex := bs(intervals, left, 0)
	right := newInterval[1]
	rightIndex := bs(intervals, right, 1) - 1
	mergeStart := 0
	mergedNewInterval := make([]int, 2)
	if leftIndex > 0 && left <= intervals[leftIndex-1][1] {
		mergeStart = leftIndex - 1
		mergedNewInterval[0] = intervals[mergeStart][0]
	} else {
		mergeStart = leftIndex
		mergedNewInterval[0] = newInterval[0]
	}

	mergeEnd := 0
	if rightIndex < len(intervals)-1 && right >= intervals[rightIndex+1][0] {
		mergeEnd = rightIndex + 1
		mergedNewInterval[1] = intervals[mergeEnd][1]
	} else {
		mergeEnd = rightIndex
		mergedNewInterval[1] = newInterval[1]
	}
	//fmt.Printf("  new=> left:%d, right:%d\n", mergeStart, mergeEnd)

	temp := make([][]int, len(intervals[:mergeStart]))
	copy(temp, intervals[:mergeStart])
	newArr := append(temp, mergedNewInterval)
	newArr = append(newArr, intervals[mergeEnd+1:]...)

	return newArr
}

func bs(intervals [][]int, v, index int) int {
	left, right := 0, len(intervals)-1
	for left <= right {
		mid := (left + right) / 2
		if intervals[mid][index] <= v {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {
	//intervals, newInterval := [][]int{{1, 3}, {6, 9}}, []int{2, 5}
	//intervals, newInterval := [][]int{{1, 5}}, []int{0, 0}
	intervals, newInterval := [][]int{{1, 5}}, []int{0, 0}
	//intervals, newInterval := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}
	//intervals, _ := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}
	//fmt.Println(bs(intervals, 8, 1) - 1)
	fmt.Println(insert(intervals, newInterval))
}
