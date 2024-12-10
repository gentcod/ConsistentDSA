package randomnote

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/ransom-note/
func Solve(ransomNote string, magazine string) {
	fmt.Println(canConstruct(ransomNote, magazine))
}

// TIME COMPLEXITY: O(n), SPACE COMPLEXITY: 0(1)

func canConstruct(ransomNote string, magazine string) bool {
	counts := make([]int, 26)
	for _, v := range magazine {
		counts[v -'a']++
	}

	for _, v := range ransomNote {
		counts[v-'a']--
		if counts[v -'a'] < 0 {
			return false
		}
	}

	return true
}

// TIME COMPLEXITY: O(n)+O(m)=O(n+m) where n is length of ransomNote, and m is length of magazine
// SPACE COMPLEXITY: 0(1)

// func canConstruct(ransomNote string, magazine string) bool {
// 	hash := make(map[rune]int)
// 	count := len(ransomNote)
// 	for _, v := range ransomNote {
// 		hash[v]++
// 	}
// 	for _,v := range magazine {
// 		if count == 0 {
// 			return true
// 		}
// 		if hash[v] > 0 {
// 			hash[v]--
// 			count--
// 		}
// 	}
// 	return count == 0
// }

