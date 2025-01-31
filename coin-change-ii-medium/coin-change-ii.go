package main

import "fmt"

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1) // dp[n][amount] = ways to gather `amount` with using last `n coins`
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j < amount+1; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(int, []int) int
	dfs = func(amount int, coins []int) int {
		if len(coins) == 1 {
			//fmt.Println(coins, amount, amount%coins[0])
			if amount%coins[0] == 0 {
				return 1
			}
			return 0
		}
		if amount == 0 {
			return 1
		} else if amount < 0 {
			return 0
		} else if dp[len(coins)][amount] != -1 {
			return dp[len(coins)][amount]
		}
		//firstAmount := amount
		ways := 0
		coin := coins[0]
		count := 0
		for amount >= 0 {
			dp[len(coins)-1][amount] = dfs(amount, coins[1:])
			ways += dp[len(coins)-1][amount]
			count += 1
			amount -= coin
		}
		//fmt.Printf("total ways for amount:%d using last %d coins: ==%d\n", firstAmount, len(coins), ways)
		return ways
	}

	return dfs(amount, coins)
}

func main() {
	//amount, coins := 5, []int{1, 2, 5}
	amount, coins := 500, []int{1, 2, 5}
	//amount, coins := 3, []int{2}
	//amount, coins := 10, []int{10}
	//amount, coins := 0, []int{10}
	fmt.Println(change(amount, coins))
}
