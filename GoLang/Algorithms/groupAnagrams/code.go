package groupanagrams

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/group-anagrams/description/
func Solve(words []string) {
	fmt.Println(groupAnagrams(words))
}

type Key [26]byte

func groupAnagrams(strs []string) [][]string {
	results := [][]string {}

	// Create an hashmap with the total ascii key as the key and an array of strings as the value
	anagrams := make(map[Key] []string)

	// Append the str array in hashmap with the value in the words array using the specific key
	for _, val := range strs {
		key := getStrKey(val)
		anagrams[key] = append(anagrams[key], val)
	}
	
	// Iterate through the map, and append the values to the results array
	for _, res := range anagrams {
		results = append(results, res)
	}

	return results
}

func getStrKey(s string) (key Key){
	for i := range s {
		key[s[i] - 'a']++
	}
	return key
}

// USING RECURSSION
// func groupAnagrams(words []string) [][]string {
// 	results := [][]string {}
// 	cur := 0
// 	count := len(words)
// 	stack := []string {}
// 	left := []string {}

// 	if count == 0 {
// 		return results
// 	}

// 	for i := 0; i < len(words); i++ {

// 		if cur == i {
// 			stack = append(stack, words[i])
// 		} else {
// 			valid := isAnagram(stack[0], words[i])
// 			if valid {
// 				stack = append(stack, words[i])
// 			} else {
// 				left = append(left, words[i])
// 			}
// 			cur++
// 		}
// 	}
// 	results = append(results, stack)
// 	count -= len(stack)
// 	fmt.Println(left, count)

// 	return append(results, groupAnagrams(left)...)
// }

// func isAnagram(s string, t string) bool {
// 	chars := make([]int, 26)

// 	for _, v := range s {
// 		i := int(v - 'a')
// 		chars[i]++
// 	}

// 	for _, v := range t {
// 		i := int(v - 'a')
// 		chars[i]--
// 	}

// 	for _, v := range chars {
// 		if v != 0 {
// 			return false
// 		}
// 	}

// 	return true
// }