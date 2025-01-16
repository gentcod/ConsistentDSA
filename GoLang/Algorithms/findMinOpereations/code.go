package findminopereations

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/
func Solve(nums []int) {
	fmt.Println(minimumOperations(nums))
}

func minimumOperations(nums []int) (c int) {
	for i := range nums {
		 if nums[i] % 3 != 0 {
			  c++
		 }
	}
	return
}