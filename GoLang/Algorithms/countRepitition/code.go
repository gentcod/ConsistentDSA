package countrepitition

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/count-the-repetitions/?source=submission-noac
func Solve(s1 string, n1 int, s2 string, n2 int) {
	fmt.Println(getMaxRepetitions(s1, n1, s2, n2))
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}

	indexr := make([]int, n1)
	countr := make([]int, n1)

	index, count := 0, 0

	for i := 0; i < n1; i++ {
		for j := 0; j < len(s1); j++ {
			if s1[j] == s2[index] {
				index++
				if index == len(s2) {
					index = 0
					count++
				}
			}
		}

		countr[i] = count
		indexr[i] = index

		for k := 0; k < i; k++ {
			if indexr[k] == index {
				prevCount := countr[k]
				patternLen := i - k
				remainingBlocks := n1 - 1 - k
				patternCount := (countr[i] - countr[k]) * (remainingBlocks / patternLen)
				remainingCount := countr[k+(remainingBlocks%patternLen)] - countr[k]

				return (prevCount + patternCount + remainingCount) / n2
			}
		}
	}

	return countr[n1-1] / n2
}

// func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) (res int) {
// 	if n2 > n1 || len(s2)*n2 > len(s1)*n1 {
// 		return 0
// 	}
// 	if s1 == s2 {
// 		if n1 == n2 {
// 			return 1
// 		} else {
// 			return n1 - n2
// 		}
// 	}

// 	if checkUnique(s1) && checkUnique(s2) {
// 		cal := (len(s1) * n1) / (len(s2) * n2)
// 		return cal
// 	}

// 	cur := s1
// 	for i := 0; i < n1; i++ {
// 		c := checkIteration(cur, s2)
// 		if c > 1 {
// 			tot := (c * n1) / n2
// 			if i > 0 {
// 				return tot / (i + 1)
// 			}
// 			return tot
// 		}

// 		if i+1 == n1 {
// 			return c
// 		}
// 		cur += s1
// 	}

// 	return
// }

// // checkUnique checks if the characters in string is a single character
// // e.g s = "aaa" or s = "aa" will return true s = "ab" will return false
// func checkUnique(s string) bool {
// 	for i := range s {
// 		if s[0] != s[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func checkIteration(s1, s2 string) int {
// 	cur, c := s2, 0
// 	for i := range s1 {
// 		if s1[i] == cur[0] {
// 			cur = cur[1:]
// 		}

// 		if len(cur) == 0 {
// 			c++
// 			cur = s2
// 		}
// 	}
// 	return c
// }
