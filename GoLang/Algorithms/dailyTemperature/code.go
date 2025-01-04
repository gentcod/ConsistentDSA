package dailytemperatures

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/daily-temperatures/
func Solve(temp []int) {
	fmt.Println(dailyTemperatures(temp))
}

func dailyTemperatures(temp []int) ([]int) {
	res := make([]int, len(temp))
	stack := []int{}

	for i := 0; i < len(temp); i++ {
		 for len(stack) > 0 && temp[stack[len(stack)-1]] < temp[i] {
			  index := stack[len(stack)-1]
			  stack = stack[:len(stack)-1]
			  res[index] = i - index
		 }
		 stack = append(stack, i)
	}

	return res
}
