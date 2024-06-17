package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		//fmt.Println("left:", nums[left], " right:", nums[right])
		if nums[left] < nums[right] {
			return nums[left]
		} else if nums[left] > nums[right] {
			mid := (right + left) / 2
			//fmt.Println("mid:", nums[mid])
			if nums[mid] < nums[right] {
				right = mid
			} else if nums[mid] > nums[right] {
				left = mid + 1
			}
		}
	}
	//fmt.Println("last -> left:", nums[left], " right:", nums[right])
	return nums[right]
}

func main() {
	//nums := []int{3, 4, 5, 1, 2}
	//nums := []int{5, 1, 2, 3, 4}
	//nums := []int{4, 5, 6, 7, 0, 1, 2, 3}
	//nums := []int{11, 13, 15, 17}
	nums := []int{17, 11, 13, 15}
	fmt.Println(findMin(nums))
}
