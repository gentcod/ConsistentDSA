package removeDuplicates

import (
	"fmt"
	// "sort"
)

// Leetcode Problem: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func Solve(arr []int) {
	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	// sort.Ints(nums)
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}