package main

import (
	"fmt"
)

func numDecodings(s string) int {

	m := map[string]bool{
		"1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true,
		"9": true, "10": true, "11": true, "12": true, "13": true, "14": true, "15": true, "16": true, "17": true,
		"18": true, "19": true, "20": true, "21": true, "22": true, "23": true, "24": true, "25": true, "26": true,
	}

	dp := make([]int, len(s)+2)

	if selfValid(s, len(s)-1, m) {
		dp[len(s)-1] = 1
	}
	if selfValid(s, len(s)-2, m) {
		if dp[len(s)-1] == 1 {
			dp[len(s)-2] = 1
		}
		if doubleValid(s, len(s)-2, m) {
			dp[len(s)-2] += 1

		}
	}
	for i := len(s) - 3; i >= 0; i-- {
		if selfValid(s, i, m) {
			dp[i] += dp[i+1]
		}
		if doubleValid(s, i, m) {
			dp[i] += dp[i+2]
		}

	}

	return dp[0]
}

func selfValid(s string, index int, m map[string]bool) bool {
	if index >= len(s) || index < 0 {
		return false
	}
	if _, exist := m[string(s[index])]; exist {
		return true
	}
	return false
}

func doubleValid(s string, index int, m map[string]bool) bool {
	if index >= len(s)-1 || index < 0 {
		return false
	}
	if _, exist := m[s[index:index+2]]; exist {
		return true
	}

	return false
}

func main() {
	//s := "12"
	//s := "226"
	//s := "06"
	s := "27"

	fmt.Println(numDecodings(s))
}
