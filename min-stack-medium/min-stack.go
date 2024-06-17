package main

import "math"

type MinStack struct {
	arr        [][2]int
	currentMin int
}

func Constructor() MinStack {
	return MinStack{arr: make([][2]int, 0), currentMin: math.MaxInt}
}

func (ms *MinStack) Push(val int) {
	ms.currentMin = min(ms.currentMin, val)
	ms.arr = append(ms.arr, [2]int{val, ms.currentMin})
}

func (ms *MinStack) Pop() {
	ms.arr = ms.arr[:len(ms.arr)-1]
	if len(ms.arr) > 0 {
		ms.currentMin = ms.arr[len(ms.arr)-1][1]
	} else {
		ms.currentMin = math.MaxInt
	}

}

func (ms *MinStack) Top() int {
	return ms.arr[len(ms.arr)-1][0]
}

func (ms *MinStack) GetMin() int {
	return ms.arr[len(ms.arr)-1][1]
}

func main() {

}
