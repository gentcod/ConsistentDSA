package countrepeatingcharacter

import (
	"fmt"
)

// Leetcode Problem: 
func Solve(s string) {
	fmt.Println(countRepeatingCharacter(s))
}

func countRepeatingCharacter(s string) (res string) {
	l, c := 0,0

	for i,char := range s {
		if string(char) != string(s[l]) || c >= 9 {
			res += fmt.Sprintf("%d%s", c, string(s[l]))
			l = i
			c = 1
		} else {
			c++
		}

		if i == len(s) - 1 {
			res += fmt.Sprintf("%d%s", c, string(char))
		}
	}
	return
}
