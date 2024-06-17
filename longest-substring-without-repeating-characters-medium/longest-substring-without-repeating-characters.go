package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	left := 0
	seen := make(map[rune]bool)
	maxWindow := 0
	for right := 0; right < len(s); right++ {
		c := rune(s[right])
		if _, ok := seen[c]; ok {
			for left < right {
				if rune(s[left]) == c {
					left += 1
					break
				} else {
					delete(seen, rune(s[left]))
				}
				left += 1
			}
		} else {
			seen[c] = true
		}
		maxWindow = max(maxWindow, right-left+1)
		//fmt.Printf("right-to-left:%v, maxWindow:%d\n", s[left:right+1], maxWindow)
	}

	return maxWindow
}

func main() {
	//s := "abcabcbb"
	//s := "bbbbb"
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
