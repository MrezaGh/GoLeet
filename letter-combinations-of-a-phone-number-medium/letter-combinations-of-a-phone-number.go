package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	numToChar := map[byte][]string{
		'2': {"a", "b", "c"}, '3': {"d", "e", "f"}, '4': {"g", "h", "i"}, '5': {"j", "k", "l"},
		'6': {"m", "n", "o"}, '7': {"p", "q", "r", "s"}, '8': {"t", "u", "v"}, '9': {"w", "x", "y", "z"}}

	var backtrack func(digits string) []string
	backtrack = func(digits string) []string {
		if len(digits) == 0 {
			return []string{}
		}

		combs := backtrack(digits[1:])
		if len(combs) == 0 {
			combs = []string{""}
		}
		newCombs := make([]string, 0)
		chars, _ := numToChar[digits[0]]
		for _, char := range chars {
			for _, comb := range combs {
				newCombs = append(newCombs, string(char)+comb)
			}
		}

		return newCombs
	}

	ans := backtrack(digits)
	return ans
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
