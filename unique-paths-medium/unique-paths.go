package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i-1][j], 1) + max(dp[i][j-1], 1)
		}
	}
	return max(1, dp[m-1][n-1])
}

func main() {
	//m, n := 3, 7
	m, n := 2, 3

	fmt.Println(uniquePaths(m, n))
}
