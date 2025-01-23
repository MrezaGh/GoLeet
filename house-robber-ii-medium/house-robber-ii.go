package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var (
		lastHomeRobbed, lastHomeSpared int
	)
	// start Robbed
	lastHomeSpared, lastHomeRobbed = -1001, nums[0]
	for i := 1; i < len(nums); i++ {
		currentRobbed, currentSpared := lastHomeSpared+nums[i], max(lastHomeRobbed, lastHomeSpared)
		lastHomeRobbed, lastHomeSpared = currentRobbed, currentSpared
	}
	maxWithRobbed := lastHomeSpared

	// start spared
	lastHomeSpared, lastHomeRobbed = 0, -1001
	for i := 1; i < len(nums); i++ {
		currentRobbed, currentSpared := lastHomeSpared+nums[i], max(lastHomeRobbed, lastHomeSpared)
		lastHomeRobbed, lastHomeSpared = currentRobbed, currentSpared
	}
	maxWithSpared := max(lastHomeRobbed, lastHomeSpared)

	return max(maxWithSpared, maxWithRobbed)
}

func main() {
	//nums := []int{2, 3, 2}
	//nums := []int{1, 2, 3, 1}
	nums := []int{1, 2, 3}

	fmt.Println(rob(nums))
}
