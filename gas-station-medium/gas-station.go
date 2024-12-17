package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	gain := make([]int, len(gas))

	for i := 0; i < len(gas); i++ {
		gain[i] = gas[i] - cost[i]
	}

	currentGain := gain[len(gain)-1]
	maxGain := currentGain
	maxIndex := len(gain) - 1

	for j := len(gain) - 2; j >= 0; j-- {
		currentGain += gain[j]
		if currentGain > maxGain {
			maxGain = currentGain
			maxIndex = j
		}
	}
	if currentGain < 0 {
		return -1
	}

	return maxIndex

}
func main() {
	gas, cost := []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}
	//gas, cost := []int{2, 3, 4}, []int{3, 4, 3}

	fmt.Println(canCompleteCircuit(gas, cost))
}
