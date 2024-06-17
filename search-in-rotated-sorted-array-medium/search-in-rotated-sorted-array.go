package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		fmt.Println("left:", nums[left], " right:", nums[right], " mid:", nums[mid])
		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[left] && nums[mid] <= nums[right] {
			//I
			if target < nums[mid] {
				right = mid - 1
			} else if target > nums[mid] {
				if target <= nums[right] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		} else if nums[mid] >= nums[left] && nums[mid] > nums[right] {
			//II
			if target > nums[mid] {
				left = mid + 1
			} else if target < nums[mid] {
				if target >= nums[left] {
					right = mid - 1
				} else if target < nums[left] {
					left = mid + 1
				}
			}
		} else {
			//III
			if target < nums[mid] {
				right = mid - 1
			} else if target > nums[mid] {
				left = mid + 1
			}
		}
	}
	fmt.Println("last -> left:", nums[left], " right:", nums[right])
	if nums[left] == target {
		return left
	}
	return -1
}
func main() {
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	//target := 0
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	//target := 3
	//nums := []int{5, 1, 3}
	//target := 5
	//nums := []int{4, 5, 6, 7, 8, 1, 2, 3}
	//target := 8
	nums := []int{3, 1}
	target := 1
	fmt.Println(search(nums, target))

}
