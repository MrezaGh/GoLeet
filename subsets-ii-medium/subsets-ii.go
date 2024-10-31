package main

import (
	"fmt"
	"slices"
)

//func subsetsWithDup(nums []int) [][]int {
//	visited := make(map[string]bool)
//	subs := recursiveSubs(nums)
//
//	//fmt.Println("before: ", subs)
//	var ans [][]int
//	for _, sub := range subs {
//		temp := make([]int, len(sub))
//		copy(temp, sub)
//		slices.Sort(temp)
//		key := fmt.Sprint(temp)
//		if _, ok := visited[key]; ok {
//			continue
//		}
//		visited[key] = true
//		ans = append(ans, temp)
//	}
//	return ans
//}
//
//func recursiveSubs(nums []int) [][]int {
//	if len(nums) == 1 {
//		return [][]int{{nums[0]}, make([]int, 0)}
//	}
//	subs := recursiveSubs(nums[1:])
//	subsWithCurrentElement := make([][]int, len(subs))
//
//	for i, sub := range subs {
//		temp := make([]int, len(sub))
//		copy(temp, sub)
//		subsWithCurrentElement[i] = append(temp, nums[0])
//	}
//	ans := make([][]int, 2*len(subs))
//	for i := 0; i < len(subs); i++ {
//		ans[i] = subs[i]
//		ans[i+len(subs)] = subsWithCurrentElement[i]
//	}
//	return ans
//}

func subsetsWithDup(nums []int) [][]int {
	subsets := make([][]int, 0)
	current := make([]int, 0)

	slices.Sort(nums)
	var backtrack func(index int)
	backtrack = func(index int) {
		subsets = append(subsets, append([]int{}, current...))
		if index == len(nums) {
			return
		}

		for i := index; i < len(nums); i++ {
			if i != index && nums[i] == nums[i-1] {
				continue
			}
			current = append(current, nums[i])
			backtrack(i + 1)
			current = current[:len(current)-1]
		}
	}
	backtrack(0)

	return subsets
}

//func subsetsWithDup(nums []int) [][]int {
//	ans := make([][]int, 0)
//	curr := make([]int, 0)
//	var backtrack func(idx int)
//	backtrack = func(idx int) {
//		ans = append(ans, append([]int{}, curr...))
//		if idx == len(nums) {
//			return
//		}
//		for i := idx; i < len(nums); i++ {
//			curr = append(curr, nums[i])
//			backtrack(i + 1)
//			curr = curr[:len(curr)-1]
//		}
//	}
//	backtrack(0)
//	return ans
//}

func main() {
	//nums := []int{1, 2, 2}
	nums := []int{1, 2, 2, 3}
	//nums := []int{9, 0, 3, 5, 7}
	//nums := []int{0}

	fmt.Println(subsetsWithDup(nums))
}
