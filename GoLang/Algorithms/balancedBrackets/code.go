package balancedbrackets

import (
	"fmt"
	"strings"
)

func Solve(str string) {
	fmt.Println(isBracketsBalanced(str))
}

// Using HashMap
func isBracketsBalanced(s string) bool {
	if s == "" {
		return false
	}
	bMap := map[string]string{
		"{": "}", "(": ")", "[": "]",
	}
	stack := []string{}
	arr := strings.Split(s, "")

	for i := 0; i < len(arr); i++ {
		if bMap[arr[i]] != "" {
			stack = append(stack, arr[i])
		} else {
			if len(stack) == 0 || bMap[stack[len(stack)-1]] != arr[i] {
				return false
			}
			stack = append(stack[:(len(stack)-1)], stack[(len(stack)-1)+1:]...)
			fmt.Println("stack end ->", stack)
		}
	}
	return len(stack) == 0
}

// Using custom func
// func isBracketsBalanced(s string) bool {

// 	stack := []string {}
// 	arr := strings.Split(s, "")

// 	for i:= 0; i < len(arr); i++ {
// 		if len(isReversed(arr[i])) != 0 {
// 			stack = append(stack, arr[i])
// 		} else {
// 			if len(stack) == 0 {
// 				return false
// 			}

// 			cur := stack[len(stack) - 1]
// 			last := len(stack) -1
// 			stack = append(stack[:last], stack[last + 1:]...)
// 			if isReversed(cur) != arr[i] {
// 				return false
// 			}
// 		}
// 	}
// 	return len(stack) == 0
// }

// func isReversed(str string) string {
// 	var reversed string
// 	if str == "{" {
// 		reversed = "}"
// 	}

// 	if str == "(" {
// 		reversed = ")"
// 	}

// 	if str == "[" {
// 		reversed = "]"
// 	}

// 	return reversed
// }

// TODO: use map
