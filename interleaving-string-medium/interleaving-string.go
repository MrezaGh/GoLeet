package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}

	cache := make(map[[3]int]bool)
	var dfs func(i, j, r int) bool
	dfs = func(i, j, r int) bool {
		key := [3]int{i, j, r}
		if v, exist := cache[key]; exist {
			return v
		} else if r == len(s3) {
			return true
		}
		res := false

		if (i < len(s1) && j < len(s2)) && s3[r] == s2[j] && s3[r] == s1[i] {
			res = dfs(i+1, j, r+1) || dfs(i, j+1, r+1)
		} else if i < len(s1) && s3[r] == s1[i] {
			res = dfs(i+1, j, r+1)
		} else if j < len(s2) && s3[r] == s2[j] {
			res = dfs(i, j+1, r+1)
		}
		cache[key] = res
		return res
	}

	return dfs(0, 0, 0)
}
func main() {
	s1, s2, s3 := "aabcc", "dbbca", "aadbbcbcac"
	//s1, s2, s3 := "aabcc", "dbbca", "aadbbbaccc"
	fmt.Println(isInterleave(s1, s2, s3))

}
