package main

import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {
	twoStepBeforeCost := 0
	oneStepBeforeCost := 0
	current := 0
	for i := 2; i <= len(cost); i++ {
		current = min(cost[i-1]+oneStepBeforeCost, cost[i-2]+twoStepBeforeCost)
		twoStepBeforeCost, oneStepBeforeCost = oneStepBeforeCost, current
		//fmt.Println(i, current)

	}

	return current
}
func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}
