package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	word1 = "^" + word1
	word2 = "^" + word2

	dp := make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		dp[i] = make([]int, len(word2))
	}
	for i := 1; i < len(word2); i++ {
		dp[0][i] = dp[0][i-1] + 1
	}
	for i := 1; i < len(word1); i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for i := 1; i < len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				fromEdit := dp[i-1][j-1] + 1
				fromInsert := dp[i][j-1] + 1
				fromDelete := dp[i-1][j] + 1
				dp[i][j] = min(fromEdit, fromInsert, fromDelete)
			}
		}
	}

	return dp[len(word1)-1][len(word2)-1]

}
func main() {
	word1, word2 := "horse", "ros"
	//word1, word2 := "intention", "execution"
	//word1, word2 := "sea", "eat"

	fmt.Println(minDistance(word1, word2))

}
