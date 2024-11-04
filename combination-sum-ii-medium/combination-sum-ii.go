package main

import (
	"fmt"
	"slices"
)

func combinationSum2(candidates []int, target int) [][]int {

	sets := make([][]int, 0)
	current := make([]int, 0)

	slices.Sort(candidates)

	var backtrack func(index, sum int)
	backtrack = func(index, sum int) {
		if sum == target {
			sets = append(sets, append([]int{}, current...))
			return
		} else if sum > target {
			return
		}

		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			current = append(current, candidates[i])
			backtrack(i+1, sum+candidates[i])
			current = current[:len(current)-1]
		}

	}

	backtrack(0, 0)
	return sets
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	fmt.Println(combinationSum2(candidates, target))
}
