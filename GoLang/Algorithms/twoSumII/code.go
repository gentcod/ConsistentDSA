package twosumII

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func Solve(numbers []int, target int) {
	fmt.Println(twoSum(numbers,target))
}

func twoSum(numbers []int, target int) (res []int) {
   hash := make(map[int]int)

	for i,n := range numbers {
		diff := target - n
		num, ok := hash[diff]

		if ok {
			res = append(res, (num + 1), i+1)
			return
		}
		hash[n] = i
	}
	return
}
