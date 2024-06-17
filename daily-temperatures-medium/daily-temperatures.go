package main

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
	var stack [][2]int

	W8 := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			stack = append(stack, [2]int{temperatures[i], i})
			W8[i] = 0
		} else if temperatures[i] < stack[len(stack)-1][0] {
			stack = append(stack, [2]int{temperatures[i], i})
			W8[i] = 1
		} else if top := stack[len(stack)-1]; temperatures[i] >= top[0] {
			for {
				if len(stack) == 0 {
					W8[i] = 0
					stack = append(stack, [2]int{temperatures[i], i})
					break
				}

				top = stack[len(stack)-1]
				if temperatures[i] < top[0] {
					stack = append(stack, [2]int{temperatures[i], i})
					W8[i] += 1
					break
				}
				stack = stack[:len(stack)-1]
				W8[i] += W8[top[1]]

			}

		}
	}
	return W8
}

func main() {
	//temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	temperatures := []int{30, 40, 50, 60}
	fmt.Println(dailyTemperatures(temperatures))
}
