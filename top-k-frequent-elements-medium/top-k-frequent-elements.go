package main

import (
	"cmp"
	"fmt"
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]int)
	freqArr := make([][]int, 0)
	for _, num := range nums {
		freqs[num] += 1
	}
	for k, v := range freqs {
		freqArr = append(freqArr, []int{k, v})
	}

	slices.SortFunc(freqArr, func(a, b []int) int {
		return cmp.Compare(b[1], a[1])
	})
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = freqArr[i][0]
	}
	//fmt.Println(freqArr)
	return ans
}

func main() {
	//nums, k := []int{1, 1, 1, 2, 2, 3}, 2
	nums, k := []int{1}, 1
	fmt.Println(topKFrequent(nums, k))
}
