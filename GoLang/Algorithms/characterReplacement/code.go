package characterreplacement

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/longest-repeating-character-replacement/
func Solve(s string, k int) {
	fmt.Println(characterReplacement(s, k))
}

func characterReplacement(s string, k int) (c int) {
	maxCount := 0
	freq := make(map[byte]int)
	l := 0

	for r := 0; r < len(s); r++ {
		freq[s[r]]++
		if freq[s[r]] > maxCount {
			maxCount = freq[s[r]]
		}

		if (r-l+1)-maxCount > k {
			freq[s[l]]--
			l++
		}

		if r-l+1 > c {
			c = r - l + 1
		}
	}

	return
}

// TO BE REVISITED
// func characterReplacement(s string, k int) (c int) {
// 	//Use an hashmap to get unique characters and how many times they appear
// 	hash := make(map[string]int)
// 	l, r, kr := 0, 0, k

// 	for r <= len(s)-1 {
// 		p, ok := hash[string(s[r])]

// 		if string(s[l]) != string(s[r]) {
// 			if kr > 0 {
// 				kr--
// 			} else {
// 				if ok && l <= p {
// 					l, r, kr = p, p, k
// 				} else {
// 					l = r
// 				}
// 			}
// 		}

// 		fmt.Println("l ->", l, "r ->", r, "cur ->", string(s[l]), "char ->", string(s[r]), kr)
// 		fmt.Println()

// 		hash[string(s[r])] = r
// 		r++
// 		if r-l > c {
// 			c = r - l
// 		}

// 	}

// 	fmt.Println(hash, s)
// 	return
// }
