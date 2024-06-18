package main

import (
	"fmt"
)

type FirstLess struct {
	left, right int
}

type MonoIncStack [][2]int

func (s *MonoIncStack) pushAndGetFirstLessIndex(num [2]int) int {
	for len(*s) > 0 {
		if num[0] > (*s)[len(*s)-1][0] {
			firstLessIndex := (*s)[len(*s)-1][1]
			*s = append(*s, num)
			return firstLessIndex
		} else {
			*s = (*s)[:len(*s)-1]
		}
	}

	if len(*s) == 0 {
		*s = append(*s, num)
		return num[1]
	}
	return 0

}

func largestRectangleArea(heights []int) int {
	leftIncStack := &MonoIncStack{}
	rightIncStack := &MonoIncStack{}
	firstLess := make([]FirstLess, len(heights))

	for i, height := range heights {
		index := leftIncStack.pushAndGetFirstLessIndex([2]int{height, i})
		firstLess[i].left = index
		if firstLess[i].left == i {
			firstLess[i].left = -1
		}
	}
	for i := len(heights) - 1; i >= 0; i-- {
		firstLess[i].right = rightIncStack.pushAndGetFirstLessIndex([2]int{heights[i], i})
		if firstLess[i].right == i {
			firstLess[i].right = len(heights)
		}
	}

	for i := 0; i < len(heights); i++ {
		fmt.Print("value:", heights[i], "  ")
		//fmt.Print("left:", "val:", heights[firstLess[i].left], "-index:", firstLess[i].left, "==")
		fmt.Print("left:", "-index:", firstLess[i].left, "==")
		//fmt.Print("right:", "val:", heights[firstLess[i].right], "-index:", firstLess[i].right, "\n")
		fmt.Print("right:", "-index:", firstLess[i].right, "==")
		fmt.Print("area:", (firstLess[i].right-firstLess[i].left-1)*heights[i], "\n")

	}
	maxArea := 0
	for i, height := range heights {
		maxArea = max(maxArea, (firstLess[i].right-firstLess[i].left-1)*height)

	}
	return maxArea
}

func main() {
	//heights := []int{2, 1, 5, 6, 2, 3}
	heights := []int{2, 4}

	fmt.Println(largestRectangleArea(heights))
}
