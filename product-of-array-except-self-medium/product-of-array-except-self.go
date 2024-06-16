package main

import "fmt"

func productExceptSelf(nums []int) []int {
	prefixProduct := make([]int, len(nums))
	postfixProduct := make([]int, len(nums))

	prefixProduct[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixProduct[i] = nums[i] * prefixProduct[i-1]
	}

	postfixProduct[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		postfixProduct[i] = postfixProduct[i+1] * nums[i]
	}
	//fmt.Println(prefixProduct, postfixProduct)

	ans := make([]int, len(nums))
	ans[0] = postfixProduct[1]
	ans[len(nums)-1] = prefixProduct[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		ans[i] = prefixProduct[i-1] * postfixProduct[i+1]
	}
	return ans
}

func main() {
	//nums := []int{1, 2, 3, 4}
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}
