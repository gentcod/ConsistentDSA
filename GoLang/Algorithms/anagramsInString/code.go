package anagramsinstring;

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/find-all-anagrams-in-a-string/
func Solve(s string, t string) {
	fmt.Println(findAnagrams(s, t))
}

func findAnagrams(s string, p string) (res []int) {
	l := 0
	r := len(p)

	for r <= len(s) {
		if getStrKey(s[l:r]) == getStrKey(p) {
			res = append(res, l)
		}
		l++
		r++
	}
	return
}

// Using Byte Signature of strings
type Key [26]byte

func getStrKey(s string) (key Key){
	for i := range s {
		key[s[i] - 'a']++
	}
	return key
}