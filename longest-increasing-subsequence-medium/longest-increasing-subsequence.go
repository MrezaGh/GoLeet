package main

import (
	"fmt"
	"slices"
)

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		innerDp := make([]int, len(nums))
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				innerDp[j] = dp[j] + 1
			} else {
				innerDp[j] = 1
			}
		}
		dp[i] = max(1, slices.Max(innerDp))
	}
	return slices.Max(dp)
}

func main() {

	//nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//nums := []int{0, 1, 0, 3, 2, 3}
	nums := []int{7, 7, 7, 7, 7, 7, 7}

	fmt.Println(lengthOfLIS(nums))

}
