package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, t := range tokens {
		switch token := string(t); token {
		case "+":
			op1, op2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, op1+op2)
		case "-":
			op1, op2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, op2-op1)
		case "*":
			op1, op2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, op1*op2)
		case "/":
			op1, op2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, op2/op1)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func main() {
	//tokens := []string{"2", "1", "+", "3", "*"}
	//tokens := []string{"4", "13", "5", "/", "+"}
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}
