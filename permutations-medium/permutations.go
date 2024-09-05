package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	permutations := make([][]int, 0)
	currPerm := make([]int, len(nums))
	visited := make(map[int]bool)

	backtrack(0, nums, currPerm, &permutations, visited)
	return permutations
}

func backtrack(idx int, nums []int, currPerm []int, permutations *[][]int, visited map[int]bool) {
	if idx == len(nums) {
		*permutations = append(*permutations, append([]int{}, currPerm...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if did, _ := visited[nums[i]]; !did {
			currPerm[idx] = nums[i]
			visited[nums[i]] = true
			backtrack(idx+1, nums, currPerm, permutations, visited)
			visited[nums[i]] = false
		}
	}
}

func main() {

	nums := []int{1, 2, 3}
	//nums := []int{0, 1}
	//nums := []int{1}
	fmt.Println(permute(nums))
}
