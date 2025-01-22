package majorityelement

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/majority-element/
func Solve(nums []int) {
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	hash := make(map[int]int)
	m,c := 0,0

	for i := range nums {
		 hash[nums[i]]++
	}

	for i,val := range hash {
		 if val > c {
			  c = val
			  m = i
		 }
	}

	return m
}