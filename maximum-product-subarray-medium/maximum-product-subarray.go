package main

import (
	"fmt"
	"slices"
)

func maxProduct(nums []int) int {
	positiveDp := make([]int, len(nums))
	negativeDp := make([]int, len(nums))
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] > 0 {
		positiveDp[0] = nums[0]
	} else {
		negativeDp[0] = nums[0]
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			positiveDp[i] = max(nums[i]*positiveDp[i-1], nums[i])
			negativeDp[i] = min(nums[i]*negativeDp[i-1], nums[i])
		} else if nums[i] < 0 {
			negativeDp[i] = min(nums[i]*positiveDp[i-1], nums[i])
			positiveDp[i] = max(nums[i]*negativeDp[i-1], nums[i])
		}
	}
	return slices.Max(positiveDp)
}

func main() {
	//nums := []int{2, 3, -2, 4}
	//nums := []int{-2, 0, -1}
	nums := []int{-2}
	fmt.Println(maxProduct(nums))
}
