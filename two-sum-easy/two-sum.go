package main

import "fmt"

func main() {
	//nums := []int{2, 7, 11, 15}
	//target := 9

	//nums := []int{3, 2, 4}
	//target := 6

	nums := []int{3, 3}
	target := 6

	fmt.Println(indexFinderOfSummation(nums, target))
}

func indexFinderOfSummation(nums []int, target int) []int {
	numberTableMapToIndex := make(map[int]int)
	for index, num := range nums {
		numberTableMapToIndex[num] = index
	}
	for index, number := range nums {
		remainder := target - number
		targetIndex, ok := numberTableMapToIndex[remainder]
		if ok && targetIndex != index {
			return []int{index, targetIndex}
		}
	}
	return nil
}
