package boiler

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/remove-element
func Solve(nums []int, val int) {
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) (res int) {
	for i := range nums {
		if nums[i] != val {
            nums[res] = nums[i]
		    res++
		}
	}
	return res
}