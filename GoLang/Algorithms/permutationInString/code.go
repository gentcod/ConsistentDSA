package permutationinstring

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/permutation-in-string/description/
func Solve(s1 string, s2 string) {
	fmt.Println(checkInclusion(s1, s2))
}

// TIME COMPLEXITY: O(n*m)
func checkInclusion(s1 string, s2 string) bool {
	bound := len(s1)
	s1keyStr := getStrKey(s1)
	if s1 == s2 {
		return true
	}
	for i := 0; i <= len(s2)-len(s1); i++ {
		if getStrKey(s2[i:bound]) == s1keyStr {
			return true
		}
		bound++
	}

	return false
}

type Key [26]byte
func getStrKey(s string) (key Key) {
	for i := range s {
		key[s[i]-'a']++
	}
	return key	
}
