package main

import "fmt"

func characterReplacement(s string, k int) int {
	biggestWindow := k
	for c := 'A'; c <= 'Z'; c++ {
		biggestWindow = max(windowFinder(s, k, c), biggestWindow)
	}
	//fmt.Println(windowFinder(s, k, 'A'))

	return biggestWindow
}

func windowFinder(s string, k int, c rune) int {
	left := 0
	currentK := k
	window := 0
	for right := 0; right < len(s); {
		currentC := rune(s[right])
		//fmt.Println(string(currentC))
		// TODO: check k =0
		if currentC == c {
			right += 1
		} else if currentK > 0 {
			right += 1
			currentK -= 1
		} else if left < right {
			for currentK <= 0 && left < right {
				if rune(s[left]) != c {
					currentK += 1
				}
				left += 1
			}
			//fmt.Println("left:", left, "right:", right)
		} else {
			left += 1
			right += 1
		}
		//fmt.Println("window:", s[left:right], "k:", currentK)
		window = max(window, right-left)

	}
	return window
}

func main() {
	//s, k := "ABAB", 2
	//s, k := "AABABBA", 1
	s, k := "AABAAABABBBA", 2
	fmt.Println(characterReplacement(s, k))
}
