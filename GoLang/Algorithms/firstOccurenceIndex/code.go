package firstoccurenceindex

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func Solve(haystack string, needle string) {
	fmt.Println(strStr(haystack,needle))
}

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if haystack == needle {
		return 0
	}

	l := -1
	for i := range haystack {
		if haystack[i] == needle[0] {
			if i+len(needle) > len(haystack) {
				return l
			}
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}

	return l
}
