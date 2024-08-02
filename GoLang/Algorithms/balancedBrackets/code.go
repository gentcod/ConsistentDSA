package balancedBrackets

import (
	"fmt"
	"strings"
)

func Solve(str string) {
	fmt.Println(isBracketsBalanced(str))
}

func isReversed(str string) (string) {
	var reversed string
	if str == "{" {
		reversed = "}"
	}

	if str == "(" {
		reversed = ")"
	}

	if str == "[" {
		reversed = "]"
	}

	return reversed
}

func isBracketsBalanced(bracketString string) bool {

	stack := []string {}
	bracketArr := strings.Split(bracketString, "")

	for i:= 0; i < len(bracketArr); i++ {
		if len(isReversed(bracketArr[i])) != 0 {
			stack = append(stack, bracketArr[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			current := stack[len(stack) - 1]
			last := len(stack) -1
			fmt.Println(stack)
			stack = append(stack[:last], stack[last + 1:]...)
			fmt.Println("->",stack)
			if isReversed(current) != bracketArr[i] {
				return false
			}
		}
	}
	return len(stack) == 0
}