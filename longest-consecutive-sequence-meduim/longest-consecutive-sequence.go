package main

import "fmt"

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)
	validsLongestLength := make(map[int]int)
	for _, num := range nums {
		numMap[num] = true
		validsLongestLength[num] = 1
	}

	for k, v := range validsLongestLength {
		longestPath := v
		next := k + 1
		for {
			if valid, ok := numMap[next]; ok {
				if valid {
					numMap[next] = false
					//longestPath += 1
					longestPath += validsLongestLength[next]
					next += 1
				} else {
					//longestPath += validsLongestLength[next]
					break
				}
			} else {
				break
			}
			//fmt.Println(longestPath)
		}
		//fmt.Printf("current longest path for %d is: %d | (valid:%v)\n", k, longestPath, numMap[k])
		validsLongestLength[k] = longestPath

	}

	maximum := 0
	for _, v := range validsLongestLength {
		maximum = max(maximum, v)
	}
	return maximum
}

func main() {
	//nums := []int{100, 4, 200, 1, 3, 2}
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	//nums := []int{5, 3, 4, 2, -4, -2, -1, -3, 0}
	fmt.Println(longestConsecutive(nums))
}
