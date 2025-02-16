package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	cache := make(map[[2]int]int)
	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		key := [2]int{i, j}
		if val, exist := cache[key]; exist {
			return val
		}
		if i == len(text1) || j == len(text2) {
			return 0
		}
		result := 0
		if text1[i] == text2[j] {
			result = dfs(i+1, j+1) + 1
		} else {
			result = max(dfs(i+1, j), dfs(i, j+1))
		}
		cache[key] = result
		return result
	}
	return dfs(0, 0)
}
func main() {
	text1, text2 := "abcde", "ace"
	//text1, text2 := "abc", "abc"
	//text1, text2 := "abc", "def"

	fmt.Println(longestCommonSubsequence(text1, text2))
}
