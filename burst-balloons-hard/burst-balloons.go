package main

import (
	"fmt"
)

func maxCoins(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}

	for length := 0; length < len(nums); length++ {
		for i := 0; i < len(nums)-length; i++ {
			j := i + length
			dp[i][j] = maxBalloonValue(dp, i, j, nums)
		}
		//fmt.Println("test:", dp)
	}

	return dp[0][len(nums)-1]
}

func maxBalloonValue(dp [][]int, i, j int, nums []int) int {
	maxValue := 0
	for index := i; index <= j; index++ {
		//fmt.Printf("i:%d, j:%d, index:%d\n", i, j, index)
		leftSide, rightSide := 0, 0
		if index-i > 0 {
			leftSide = dp[i][index-1]
		}
		if j-index > 0 {
			rightSide = dp[index+1][j]
		}
		deleteValue := popBalloon(nums, i-1, j+1, index)
		maxValue = max(maxValue, leftSide+rightSide+deleteValue)
	}
	return maxValue
}

func popBalloon(nums []int, left, right, index int) int {
	var leftItem, rightItem int
	if left < 0 {
		leftItem = 1
	} else {
		leftItem = nums[left]
	}

	if right >= len(nums) {
		rightItem = 1
	} else {
		rightItem = nums[right]
	}
	return rightItem * leftItem * nums[index]

}

func main() {
	nums := []int{3, 1, 5, 8}
	//nums := []int{1, 5}
	fmt.Println(maxCoins(nums))
}
