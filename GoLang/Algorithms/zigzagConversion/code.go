package zigzagConversion

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/zigzag-conversion/
func Solve(s string, n int) {
	fmt.Println(convert(s, n))
}

func convert(s string, n int) (res string) {
	if n <= 1 {
		return s
	}

	rows := make([]string, n)
	ns := n - 2
	rc, nsc, rcI, ncI := n, ns, 0, ns

	for i := range s {
		if rc > 0 {
			rows[rcI] += string(s[i])
			rc--
			rcI++
			if n == 2 && rc == 0 {
				rc, rcI = n, 0
			}
			continue
		}

		if rc == 0 && nsc >= 0 {
			rows[ncI] += string(s[i])
			nsc--
			ncI--
		}

		if rc == 0 && nsc == 0 {
			rc, nsc, rcI, ncI = n, ns, 0, ns
		}
	}

	for i := range rows {
		res += rows[i]
	}
	return
}
