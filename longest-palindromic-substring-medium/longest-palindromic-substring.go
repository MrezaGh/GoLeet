package main

import "fmt"

func longestPalindrome(s string) string {
	maxLen := 0
	maxPalindrome := string(s[0])
	for i := 1; i < len(s); i++ {
		evenPal := evenPalindrome(s, i)
		oddPal := oddPalindrome(s, i)
		if curMax := max(len(evenPal), len(oddPal)); curMax > maxLen {
			maxLen = curMax
			if len(evenPal) > len(oddPal) {
				maxPalindrome = evenPal
			} else {
				maxPalindrome = oddPal
			}
		}
	}

	return maxPalindrome
}

func oddPalindrome(s string, index int) string {
	i, j := index-1, index+1
	pal := string(s[index])
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		} else {
			pal = string(s[i]) + pal + string(s[j])
		}
		i--
		j++
	}

	return pal
}

func evenPalindrome(s string, index int) string {
	if s[index-1] != s[index] {
		return string(s[index])
	}
	i, j := index-2, index+1
	pal := s[index-1 : index+1]
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		} else {
			pal = string(s[i]) + pal + string(s[j])
		}
		i--
		j++
	}

	return pal
}

func main() {
	s := "babad"
	//s := "cbbd"
	//s := "d"

	fmt.Println(longestPalindrome(s))
}
