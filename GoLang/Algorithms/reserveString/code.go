package reversestring

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/reverse-string/description/
func Solve(s []string) {
	reverseString(s)
	fmt.Println(s)
}

func reverseString(s []string) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		r--
		l++
	}
}

// PURE FUNCTION: Doesn't modify input
// func reverseString(s []string) []string {
// 	res := make([]string, len(s))
// 	l, r := 0, len(s)-1

// 	for r >= 0 {
// 		res[l] = s[r]
// 		r--
// 		l++
// 	}

// 	return res
// }
