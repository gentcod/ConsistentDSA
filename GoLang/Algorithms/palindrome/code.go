package palindrome

import (
	"fmt"
	"regexp"
	"strings"
	// "unicode"
)

func Solve(val string) {
	fmt.Println(isPalindrome(val))
}

func isPalindrome(s string) bool {
	val := clean(s)
	i, j := 0, len(val)-1

	for i < j {
		if val[i] != val[j] {
			return false
		}
		i = i + 1
		j = j - 1
	}

	return true
}

func clean(str string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	if re.MatchString(str) {
		return strings.ToLower((re.ReplaceAllString(str, "")))
	}
	return str
}

// ALTERNATE SOLUTION
// func isPalindrome(s string) bool {
// 	clean := func(r rune) rune {
// 		 if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
// 			  return -1
// 		 }
// 		 return unicode.ToLower(r)
// 	}
// 	s = strings.Map(clean, s)
	
// 	i, j := 0, len(s)-1
	
// 	for i < j {
// 		 if s[i] != s[j] {
// 			  return false
// 		 }
// 		 i++
// 		 j--
// 	}
	
// 	return true
// }
