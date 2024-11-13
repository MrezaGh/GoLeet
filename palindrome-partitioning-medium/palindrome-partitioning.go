package main

import "fmt"

func partition(s string) [][]string {
	ans := make([][]string, 0)
	cur := make([]string, 0)

	var backtrack func(s string)

	backtrack = func(s string) {
		if s == "" {
			ans = append(ans, append([]string{}, cur...))
			//cur = []string{}
			return
		}
		for i := 0; i < len(s); i++ {
			if palindrome(s[:i+1]) {
				cur = append(cur, s[:i+1])
				backtrack(s[i+1:])
				cur = cur[:len(cur)-1]
			}
		}

	}
	backtrack(s)

	return ans
}

func palindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
func main() {

	s := "aaba"
	//s := "a"
	fmt.Println(partition(s))

}
