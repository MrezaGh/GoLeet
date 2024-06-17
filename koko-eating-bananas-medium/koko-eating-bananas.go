package main

import (
	"fmt"
)

func minEatingSpeed(piles []int, h int) int {
	sum := 0
	maxItem := 0
	for _, v := range piles {
		sum += v
		maxItem = max(maxItem, v)
	}
	//fmt.Println(piles)
	k := max(1, sum/h)
	k = search(piles, h, k, maxItem)
	return k
}

func spentHours(piles []int, k int) int {
	hours := 0
	for _, v := range piles {
		hours += v / k
		if v%k != 0 {
			hours += 1
		}
	}
	return hours
}

func search(piles []int, target int, k int, maxItem int) int {
	lowerBound, upperBound := k, maxItem
	for lowerBound < upperBound {
		mid := (lowerBound + upperBound) / 2
		hours := spentHours(piles, mid)
		fmt.Println("k:", mid, " spentHours:", hours, " lower:", lowerBound, "upper:", upperBound)
		if hours > target {
			lowerBound = mid + 1
		} else if hours <= target {
			upperBound = mid
		}
	}
	return upperBound

}

func main() {
	//piles := []int{3, 6, 7, 11}
	//h := 8
	//piles := []int{2, 3, 4, 5, 8}
	//h := 7
	//piles := []int{30, 11, 23, 4, 20}
	//h := 5
	//piles := []int{30, 11, 23, 4, 20}
	//h := 6
	//piles := []int{2, 3, 3, 4, 5, 5, 5, 13}
	//h := 9
	//piles := []int{2, 2}
	//h := 2
	//piles := []int{312884470}
	//h := 312884469
	piles := []int{312884470}
	h := 968709470

	//fmt.Println(search(piles, 9))
	fmt.Println(minEatingSpeed(piles, h))
}
