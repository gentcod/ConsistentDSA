package longestsubstring

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/longest-substring-without-repeating-characters/
func Solve(s string) {
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) (c int) {
	hash := make(map[rune]int)
	l,r := 0,0

	for i, char := range s {
		p, ok := hash[char]
		if ok && l <= p+1 {
			l=p+1
		} 
		
		r++
		hash[char] = i
		if r-l > c {
			c = r-l
		}

	}
	return
}


