package maxsubarray

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/maximum-subarray/
func Solve(nums []int) {
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	maxn := nums[0]
	sum := 0

	for _,n := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += n
		maxn = max(maxn, sum)
	}
	return maxn
}

