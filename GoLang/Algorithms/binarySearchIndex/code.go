package boiler

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/binary-search/
func Solve(list []int, num int) {
	fmt.Println(search(list, num))
}

func search(nums []int, target int) (res int) {
	l, r := 0, len(nums)

	for l < r {
		m := l + (r-l)/2
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	if l < len(nums) && nums[l] == target {
		return l
	}
	return -1
}
