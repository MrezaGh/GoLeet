package main

import "fmt"

func partitionLabels(s string) []int {
	lastAppearance := make(map[string]int)
	for i := 0; i < len(s); i++ {
		lastAppearance[string(s[i])] = i
	}
	//fmt.Println(lastAppearance)

	ans := make([]int, 0)
	partitionStart := 0
	partitionEnd := 0
	for i := 0; i < len(s); i++ {
		partitionEnd = max(partitionEnd, lastAppearance[string(s[i])])
		if i == partitionEnd {
			ans = append(ans, partitionEnd-partitionStart+1)
			partitionStart = i + 1
		}
	}
	return ans
}

func main() {
	//s := "ababcbacadefegdehijhklij"
	s := "eccbbbbdec"

	fmt.Println(partitionLabels(s))
}
