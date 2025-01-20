package lengthoflastword

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/length-of-last-word/
func Solve(s string) {
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	c := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if c > 0 {
				return c
			}
			continue
		}
		c++
	}
	return c
}
