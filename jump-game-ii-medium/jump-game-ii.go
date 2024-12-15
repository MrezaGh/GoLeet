package main

import (
	"fmt"
)

func jump(nums []int) int {
	maxReachable := make(map[int]int)
	currentTurn := 0
	maxReachable[0] = 0
	if len(nums) == 1 {
		return 0
	}
	for i, num := range nums {
		for i > maxReachable[currentTurn] {
			currentTurn++
		}

		reachable := i + num
		//fmt.Println("first:", i, currentTurn, reachable, maxReachable)

		maxReachable[currentTurn+1] = max(maxReachable[currentTurn+1], reachable)

		//fmt.Println(i, currentTurn, reachable, maxReachable)
		if reachable >= len(nums)-1 {
			return currentTurn + 1
		}

	}
	return -1
}

func main() {
	//nums := []int{2, 3, 1, 1, 4}
	//nums := []int{2, 3, 0, 1, 4}
	nums := []int{1}

	fmt.Println(jump(nums))

}
