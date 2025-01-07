package issubsequence

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/is-subsequence/
func Solve(s string, t string) {
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	cur := s
	for i := range t {
		if len(cur) == 0 {
			return true
		}
		if t[i] == cur[0] {
			cur = cur[1:]
		}
	}
	return len(cur) == 0
}
