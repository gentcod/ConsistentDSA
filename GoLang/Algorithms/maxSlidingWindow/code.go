package maxslidingwindow

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/sliding-window-maximum/
func Solve(nums []int, k int) {
	fmt.Println(maxSlidingWindow(nums, k))
}

func maxSlidingWindow(nums []int, k int) (res []int) {
	deque := make([]int, 0, k)

	for i, n := range nums {
		for len(deque) > 0 && nums[deque[len(deque)-1]] < n {
			deque = deque[:len(deque)-1]
		}
		fmt.Println("before: ",deque)
		deque = append(deque, i)

		if i < k-1 {
			continue
		}

		res = append(res, nums[deque[0]])

		if deque[0] == i-k+1 {
			deque = deque[1:]
		}
	}
	return
}

// WORKS BUT FAILS DUE TO TIME COMPLEXITIES
// func maxSlidingWindow(nums []int, k int) (res []int) {
// 	l := 0
// 	r := k
// 	for r <= len(nums) {
// 		res = append(res, max(nums[l:r]))
// 		l++
// 		r++
// 	}
// 	return
// }

// func max(nums []int) int {
// 	m := nums[0]
// 	for _, n := range nums {
// 		if n > m {
// 			m = n
// 		}
// 	}
// 	return m
// }
