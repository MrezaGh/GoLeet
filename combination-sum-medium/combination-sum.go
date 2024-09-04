package main

import (
	"fmt"
)

//func combinationSum(candidates []int, target int) [][]int {
//	result := combinationSumRec(candidates, target)
//	cache := make(map[string]bool)
//	prunedResult := make([][]int, 0)
//	for _, res := range result {
//		slices.Sort(res)
//		encoded := fmt.Sprint(res)
//		if _, exist := cache[encoded]; !exist {
//			prunedResult = append(prunedResult, res)
//			cache[fmt.Sprint(res)] = true
//		}
//	}
//	return prunedResult
//}
//
//func combinationSumRec(candidates []int, target int) [][]int {
//	if target < 0 {
//		return nil
//	}
//	var result [][]int
//	for _, candidate := range candidates {
//		if target-candidate == 0 {
//			result = append(result, []int{candidate})
//		} else if combinations := combinationSumRec(candidates, target-candidate); combinations != nil {
//			for i, _ := range combinations {
//				//fmt.Println("working comb:", comb, "target:", target-candidate)
//				combinations[i] = append(combinations[i], candidate)
//			}
//			//fmt.Println("combs:", combinations, target)
//			result = append(result, combinations...)
//			//fmt.Println(result)
//		}
//	}
//	if len(result) == 0 {
//		return nil
//	}
//
//	return result
//}

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	curr := make([]int, 0)

	var backtrack func(idx int, currSum int)
	backtrack = func(idx int, currSum int) {
		if currSum == target {
			newSolution := make([]int, len(curr))
			copy(newSolution, curr)
			ans = append(ans, newSolution)
			return
		} else if currSum > target {
			return
		}

		for i := idx; i < len(candidates); i++ {
			curr = append(curr, candidates[i])
			backtrack(i, currSum+candidates[i])
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0, 0)
	return ans
}

func main() {
	//candidates := []int{2, 3, 6, 7}
	//target := 7

	candidates := []int{2, 3, 5}
	target := 8

	//candidates := []int{2}
	//target := 1

	fmt.Println(combinationSum(candidates, target))
}
