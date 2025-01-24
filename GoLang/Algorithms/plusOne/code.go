package plusone

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/plus-one/
func Solve(digits []int) {
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) (res []int) {
	res = digits
	iter, rl, l := 1, len(digits), 1
	for iter > 0 {
		if rl-l == -1 {
			n := []int{}
			n = append(n, res[0] + 1)
			n = append(n, res...)
			return n
		}
		p := digits[rl-l] + 1
		if p > 9 {
			p = 10 - p
			res[rl-l] = p
			l++
			continue
		}
		res[rl-l] = p
		return
	}

	return
}
