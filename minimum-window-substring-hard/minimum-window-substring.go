package main

import (
	"fmt"
	"maps"
)

type window struct {
	length      int
	left, right int
}

func minWindow(s string, t string) string {
	var minW window = window{length: len(s) + 1}
	needed, seen := make(map[string]int), make(map[string]int)
	for _, r := range t {
		needed[string(r)] += 1
	}
	maps.Copy(seen, needed)
	left := 0
	for right := 0; right < len(s); right++ {
		c := string(s[right])
		if _, ok := needed[c]; ok {
			needed[c] -= 1
			if needed[c] == 0 {
				delete(needed, c)
			}
		}
		if _, ok := seen[c]; ok {
			seen[c] -= 1
			// push left
			for left <= right {
				lChar := string(s[left])
				if count, ok := seen[lChar]; ok && count < 0 {
					//fmt.Println("pushing left:", left, lChar, "count:", count)
					left += 1
					seen[lChar] += 1
				} else if !ok {
					//fmt.Println("pushing spam left:", left, lChar)
					left += 1
				} else {
					break
				}
			}
		}
		//fmt.Printf("window:%v, needed:%v, seen:%v\n", s[left:right+1], needed, seen)

		if len(needed) == 0 {
			newW := window{length: right - left + 1, left: left, right: right}
			if newW.length < minW.length {
				minW = newW
			}
			//fmt.Println("bingo", minW)
		}
	}

	if minW.length > len(s) {
		return ""
	} else {
		return s[minW.left : minW.right+1]
	}

}

func main() {
	//s, t := "ADOBECODEBANC", "ABC"
	s, t := "ADOOBECODEBANC", "ABDC"
	//s, t := "a", "a"
	//s, t := "a", "aa"
	fmt.Println(minWindow(s, t))

}
