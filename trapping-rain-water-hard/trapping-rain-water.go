package main

import (
	"fmt"
	"slices"
)

func trap(height []int) int {
	trappedRain := 0
	leftAsc := make([][2]int, 0)
	maxLeft := -1
	rightAsc := make([][2]int, 0)
	maxRight := -1
	for left := 0; left < len(height); left++ {
		if height[left] >= maxLeft {
			leftAsc = append(leftAsc, [2]int{height[left], left})
			maxLeft = height[left]
		}
	}
	for right := len(height) - 1; right >= 0; right-- {
		if height[right] > maxRight {
			rightAsc = append(rightAsc, [2]int{height[right], right})
			maxRight = height[right]
		}
	}
	slices.Reverse(rightAsc)
	// fmt.Println(leftAsc, rightAsc)
	for i := 0; i < len(leftAsc)-1; i++ {
		water := pot(height[leftAsc[i][1] : leftAsc[i+1][1]+1])
		// fmt.Println("found pot:", leftAsc[i][1], leftAsc[i+1][1], "-->", height[leftAsc[i][1]:leftAsc[i+1][1]+1], "water:", water)
		trappedRain += water
	}

	for i := 0; i < len(rightAsc)-1; i++ {
		water := pot(height[rightAsc[i][1] : rightAsc[i+1][1]+1])
		// fmt.Println("found pot:", rightAsc[i][1], rightAsc[i+1][1], "-->", height[rightAsc[i][1]:rightAsc[i+1][1]+1], "water:", water)
		trappedRain += water
	}
	return trappedRain
}

func pot(heights []int) int {
	sum := 0
	for i := 1; i < len(heights)-1; i++ {
		sum += heights[i]
	}

	return min(heights[0], heights[len(heights)-1])*(len(heights)-2) - sum
}
func main() {
	//height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	height := []int{4, 2, 0, 3, 2, 5}

	fmt.Println(trap(height))
	//fmt.Println(pot(height))
}
