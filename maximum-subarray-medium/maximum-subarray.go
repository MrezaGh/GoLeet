



package main

import "fmt"

//func maxSubArray(nums []int) int {
//
//	l, r := 0, 0
//	sum := nums[0]
//	maxSum := sum
//
//	for {
//		if r < len(nums)-1 && l < r {
//			shiftRightSum := sum + nums[r+1]
//			shiftLeftSum := sum - nums[l]
//			if shiftRightSum >= shiftLeftSum {
//				sum += nums[r+1]
//				r++
//			} else {
//				sum -= nums[l]
//				l++
//			}
//
//		} else if l < r {
//			sum -= nums[l]
//			l++
//		} else {
//			if r == len(nums)-1 {
//				return maxSum
//			}
//			sum += nums[r+1]
//			r++
//
//		}
//		//fmt.Printf("l:%d ,r:%d. sum:%d \n", l, r, sum)
//		maxSum = max(sum, maxSum)
//
//	}
//}

func maxSubArray(nums []int) int {
	maxSum := nums[0]

	sumFromStart := 0
	seenMinSum := 0
	for _, num := range nums {
		sumFromStart += num
		maxSum = max(maxSum, sumFromStart-seenMinSum)
		seenMinSum = min(seenMinSum, sumFromStart)

	}

	return maxSum
}
func main() {
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//nums := []int{1}
	//nums := []int{5, 4, -1, 7, 8}
	nums := []int{2, -3, 1, 3, -3, 2, 2, 1}
	fmt.Println(maxSubArray(nums))
}
