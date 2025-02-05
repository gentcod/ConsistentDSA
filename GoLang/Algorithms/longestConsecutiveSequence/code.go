package longestconsecutive

import (
	"fmt"
	"sort"
)

// Leetcode Problem: https://leetcode.com/problems/longest-consecutive-sequence/
func Solve(nums []int) {
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) (c int) {
	if len(nums) == 0 {
		return
	}

	if len(nums) == 1 {
		return 1
	}

	sort.Ints(nums)
	prev, count := nums[0], 0
	for i, n := range nums {
		if prev == n && i != 0 {
			continue
		}

		if n-prev == 1 {
			count++
		} else {
			count = 1
		}

		if count > c {
			c = count
		}
		prev = n
	}
	return
}

// ALTERNATIVE APPROACH
// func longestConsecutive(nums []int) (c int) {
// 	if len(nums) == 1 {
// 		return 1
// 	}
// 	sort.Ints(nums)
// 	set := makeSet(nums)

// 	count := 1
// 	if len(set) == 1 {
// 		return 1
// 	}
// 	for i, n := range set {
// 		if i+1 >= len(set) {
// 			return
// 		}
// 		if set[i+1]-n == 1 {
// 			count++
// 		} else {
// 			count = 1
// 		}
// 		if count > c {
// 			c = count
// 		}
// 	}
// 	return
// }

// // makeSet takes a sorted list or arrays of integers and returns a list or array with only unique integers
// func makeSet(nums []int) (set []int) {
// 	c := 0
// 	for _, n := range nums {
// 		if n == 0 {
// 			if len(set) < 1 || set[len(set)-1] < 0 {
// 				set = append(set, n)
// 			}
// 		}
// 		if n != 0 && n != c {
// 			set = append(set, n)
// 			c = n
// 		}
// 	}
// 	return
// }
