package main

import "fmt"

func numDistinct(s string, t string) int {

	dp := make(map[[2]int]int) // dp[{1, 2}] = the number of ways to match  1:end characters of s with the 2:end characters of t

	var dfs func(int, int) int
	dfs = func(SIndex int, TIndex int) int {
		if TIndex == len(t) {
			return 1
		} else if SIndex == len(s) {
			return 0
		} else if found, exist := dp[[2]int{SIndex, TIndex}]; exist {
			return found
		}

		ways := 0
		for i := SIndex; i <= len(s)-(len(t)-TIndex); i++ { // len(t)-TIndex => the remaining length of t
			if s[i] == t[TIndex] {
				found := dfs(i+1, TIndex+1)
				ways += found
				//fmt.Printf("calling dfs with s=%s, t=%s => found:%d \n", s[i+1:], t[TIndex+1:], found)

				//ways += dfs(i+1, TIndex+1)
			}
		}
		dp[[2]int{SIndex, TIndex}] = ways
		return ways
	}

	return dfs(0, 0)
}

func main() {
	s, t := "rabbbit", "rabbit"
	//s, t := "babgbag", "bag"

	fmt.Println(numDistinct(s, t))
}
