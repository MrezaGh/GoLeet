package main

import "fmt"

//func generateParenthesis(n int) []string {
//	var result []string
//	for i := 0; i < 2<<(2*n-1); i++ {
//		s := fmt.Sprintf("%0*b", 2*n, i)
//		//fmt.Println(i, s, validate(s))
//		if validate(s) {
//			result = append(result, s)
//		}
//	}
//	result = prettyfie(result)
//	return result
//}
//
//func prettyfie(result []string) []string {
//	var pretty []string
//	for _, s := range result {
//		parentheses := ""
//		for _, c := range s {
//			if c == '1' {
//				parentheses += ")"
//			} else {
//				parentheses += "("
//			}
//		}
//		pretty = append(pretty, parentheses)
//	}
//	return pretty
//}
//
//func validate(str string) bool {
//	var stack []rune
//
//	for _, c := range str {
//		if c == '1' {
//			if len(stack) == 0 {
//				return false
//			} else if stack[len(stack)-1] == '0' {
//				stack = stack[:len(stack)-1]
//			} else {
//				return false
//			}
//		} else {
//			stack = append(stack, c)
//		}
//	}
//	if len(stack) != 0 {
//		return false
//	}
//	return true
//}

// better solution
func generateParenthesis(n int) []string {
	var result []string
	result = backtrack(0, 0, "", 2*n)
	return result
}

func backtrack(countOpen int, countClose int, str string, remain int) []string {
	var result []string
	if remain == 0 {
		return []string{str}
	}
	if countOpen > countClose && remain > countOpen-countClose {
		result = append(result, backtrack(countOpen+1, countClose, str+"(", remain-1)...)
		result = append(result, backtrack(countOpen, countClose+1, str+")", remain-1)...)
	} else if countOpen > countClose && remain == countOpen-countClose {
		result = append(result, backtrack(countOpen, countClose+1, str+")", remain-1)...)
	} else if countOpen == countClose {
		result = append(result, backtrack(countOpen+1, countClose, str+"(", remain-1)...)
	}
	//fmt.Println(countOpen, countClose, str, remain)
	//fmt.Println(result)
	return result

}

func main() {
	n := 13
	//fmt.Println(generateParenthesis(n))
	fmt.Println(len(generateParenthesis(n)))
}
