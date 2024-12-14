package main

import "fmt"

func canJump(nums []int) bool {
	//reachable := make([]bool, len(nums))
	//reachable[0] = true
	maxReachable := 0
	for i, num := range nums {
		if i > maxReachable {
			return false
		} else if maxReachable >= len(nums)-1 {
			return true
		}
		maxReachable = max(i+num, maxReachable)
		//fmt.Println(num, maxReachable)

	}

	return false
}
func main() {
	//nums := []int{2, 3, 1, 1, 4}
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}
