package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	cache := make(map[[2]int]int)
	var dfs func(index, target int) int

	dfs = func(index, target int) int {
		key := [2]int{index, target}
		if val, exist := cache[key]; exist {
			return val
		} else if target == 0 && index == len(nums) {
			return 1
		} else if index == len(nums) {
			return 0
		}

		ways := dfs(index+1, target-nums[index]) + dfs(index+1, target+nums[index])
		cache[key] = ways
		return ways
	}

	return dfs(0, target)
}
func main() {
	nums, target := []int{1, 1, 1, 1, 1}, 3
	//nums, target := []int{1}, 1
	fmt.Println(findTargetSumWays(nums, target))
}
