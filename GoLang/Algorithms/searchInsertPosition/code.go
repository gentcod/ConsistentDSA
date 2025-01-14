package searchinsertposition

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/search-insert-position/
func Solve(nums []int, target int) {
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) (res int) {
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	for i := range nums {
		if target == nums[i] || target < nums[i] {
			return i
		}
	}

	return len(nums)
}
