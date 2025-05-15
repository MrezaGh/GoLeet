package main

import "fmt"

func missingNumber(nums []int) int {
	expected_sum := (len(nums) + 1) * len(nums) / 2
	var sum int
	for _, num := range nums {
		sum += num
	}
	return expected_sum - sum
}
func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
}
