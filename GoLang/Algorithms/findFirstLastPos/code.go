package findfirstlastpos

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func Solve(nums []int, target int) {
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}

	for i := range nums {
		if nums[i] == target {
			if res[0] >= 0 {
				res[1] = i
			} else {
				res[0], res[1] = i, i
			}
		}
	}

	return res
}
