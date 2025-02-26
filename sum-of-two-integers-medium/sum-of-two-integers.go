package main

import (
	"fmt"
)

func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}
func main() {
	a, b := 200, 3

	fmt.Println(getSum(a, b))
}
