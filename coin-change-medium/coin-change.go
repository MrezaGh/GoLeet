package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make(map[int]int)

	var dfs func(amount int) int
	dfs = func(amount int) int {
		if amount == 0 {
			return 0
		} else if val, exist := dp[amount]; exist {
			return val
		}

		minCoins := math.MaxInt32
		for _, coin := range coins {
			if amount-coin >= 0 {
				minCoins = min(minCoins, 1+dfs(amount-coin))
			}
		}
		dp[amount] = minCoins
		return minCoins
	}
	ans := dfs(amount)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	//coins, amount := []int{1, 2, 5}, 11
	coins, amount := []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864
	//coins, amount := []int{2}, 3
	//coins, amount := []int{1}, 0
	fmt.Println(coinChange(coins, amount))

}
