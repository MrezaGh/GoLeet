package main

import (
	"fmt"
	"slices"
)

func isNStraightHand(hand []int, groupSize int) bool {
	picked := make([]bool, len(hand))
	slices.Sort(hand)
	//fmt.Println(hand)
	for i, _ := range hand {
		if picked[i] {
			continue
		}
		if !buildGroup(i, picked, hand, groupSize) {
			return false
		}
	}
	return true
}

func buildGroup(currentIndex int, picked []bool, hand []int, size int) bool {
	seriesLength := 0
	prevIndex := currentIndex
	for i := currentIndex; seriesLength < size; i++ {
		if i >= len(hand) {
			return false
		}

		if picked[i] || (seriesLength > 0 && hand[i] == hand[prevIndex]) {
			continue
		}

		picked[i] = true
		if seriesLength == 0 {
			seriesLength += 1
		} else if hand[i] == hand[prevIndex]+1 {
			seriesLength += 1
			prevIndex = i
		} else {
			return false
		}
		//fmt.Printf("%d ", hand[i])

	}
	//fmt.Println("")
	return true
}

func main() {
	//hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	//groupSize := 3

	hand := []int{1, 2, 3, 4, 5}
	groupSize := 4

	fmt.Println(isNStraightHand(hand, groupSize))
}
