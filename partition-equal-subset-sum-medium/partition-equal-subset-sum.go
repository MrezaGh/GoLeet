package main

import (
	"fmt"
	"maps"
)

func canPartition(nums []int) bool {
	target := 0
	for _, num := range nums {
		target += num
	}
	if target%2 != 0 {
		return false
	}
	target /= 2
	possibilities := map[int]bool{nums[0]: true}
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		temp := maps.Clone(possibilities)
		for k := range maps.Keys(temp) {
			possibilities[k+num] = true
			if k+num == target {
				return true
			}
		}
	}
	if _, exist := possibilities[target]; exist {
		return true
	}
	return false
}

func main() {
	//nums := []int{1, 5, 11, 5}
	//nums := []int{1, 2, 3, 5}
	nums := []int{2, 2, 3, 5}
	fmt.Println(canPartition(nums))
}
