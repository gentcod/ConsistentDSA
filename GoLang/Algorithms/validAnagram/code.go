package validAnagram;

import (
	"fmt"
	// "reflect"
	// "sort"
	// "strings"
)

// Leetcode Problem: https://leetcode.com/problems/valid-anagram/
func Solve(s string, t string) {
	fmt.Println(validAnagram(s, t))
}

// Using Byte Signature of strings
type Key [26]byte

func validAnagram(s string, t string) bool {
	sByte := getStrKey(s)
	tByte := getStrKey(t)

	return sByte == tByte
}

func getStrKey(s string) (key Key){
	for i := range s {
		key[s[i] - 'a']++
	}
	return key
}

// Using HashMap
// func validAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	sMap := make(map[string]int)
// 	tMap := make(map[string]int)
// 	first := strings.Split(s, "")
// 	second := strings.Split(t, "")

// 	for i := 0; i < len(s); i++ {
// 		vals := sMap[first[i]]
// 		sMap[first[i]] = 1 + vals

// 		valt := tMap[second[i]]
// 		tMap[second[i]] = 1 + valt
// 	}

// 	return reflect.DeepEqual(sMap, tMap)
// }

// Using Sorting
// func validanagram(s string, t string) bool {
// 	first := strings.Split(s, "")
// 	second := strings.Split(t, "")
// 	sort.Strings(first)
// 	sort.Strings(second)
// 	return strings.Join(first, "") == strings.Join(second, "")
// }

// REF
func IsAnagram(s string, t string) bool {
	chars := make([]int, 26)

	for _, v := range s {
		i := int(v - 'a')
		chars[i]++
	}

	for _, v := range t {
		i := int(v - 'a')
		chars[i]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}

// SAMPLE DATA
// strs := []string {"eat","tea","tan","ate","nat","bat"}