package main

import (
	"fmt"
)

type status struct {
	buy   bool
	index int
}

func maxProfit(prices []int) int {
	n := len(prices)
	seen := make(map[status]int)
	var dfs func(buy bool, index int) int
	dfs = func(buy bool, index int) int {
		if index >= n {
			return 0
		} else if ans, exist := seen[status{buy, index}]; exist {
			return ans
		}
		result := 0
		if buy {
			here := dfs(false, index+1) - prices[index]
			next := dfs(true, index+1)
			result = max(here, next)
		} else {
			here := dfs(true, index+2) + prices[index]
			next := dfs(false, index+1)
			result = max(here, next)
		}
		seen[status{buy, index}] = result
		return result
	}

	maxP := dfs(true, 0)
	return maxP
}

func main() {
	//prices := []int{1, 2, 3, 0, 2}
	//prices := []int{1}
	prices := []int{2, 1, 4}

	fmt.Println(maxProfit(prices))
}
