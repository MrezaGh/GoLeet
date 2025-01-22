package main

import "fmt"

func rob(nums []int) int {
	lastHomeStolen, lastHomeSpared := 0, 0
	for _, num := range nums {
		currentStolen := lastHomeSpared + num
		currentSpared := lastHomeStolen
		lastHomeStolen, lastHomeSpared = currentStolen, currentSpared
	}
	return max(lastHomeSpared, lastHomeStolen)
}

func main() {
	//nums := []int{1, 2, 3, 1}
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}
