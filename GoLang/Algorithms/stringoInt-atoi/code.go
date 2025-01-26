package stringtointegeratoi

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Leetcode Problem: https://leetcode.com/problems/string-to-integer-atoi/
func Solve(s string) {
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "-" || s == "+" || len(s) == 0 {
		return 0
	}
	hash := map[string]bool{"-+": true, "--": true, "+-": true, "++": true}
	if len(s) > 2 {
		_, ok := hash[s[:2]]
		if ok {
			return 0
		}
	}
	if n, err := strconv.Atoi(s); err == nil {
		if n < math.MinInt32 {
			return math.MinInt32
		}
		if n > math.MaxInt32 {
			return math.MaxInt32
		}
		return n
	}

	m, l, r := 1, -1, 0
	for i := range s {
		if string(s[i]) == "-" || string(s[i]) == "+" {
			if i == 0 {
				if string(s[i]) == "-" {
					m = -1
				}
				if n, err := strconv.Atoi(s[i+1:]); err == nil {
					v := n * m
					if v < math.MinInt32 {
						return math.MinInt32
					}
					if v > math.MaxInt32 {
						return math.MaxInt32
					}
					return v
				}
				continue
			}
		}

		if _, err := strconv.Atoi(string(s[i])); err == nil {
			if l == -1 {
				l = i
			}
		} else {
			if l == -1 {
				return 0
			}
			r = i
			n, _ := strconv.Atoi(string(s[l:r]))
			v := n * m
			if v < math.MinInt32 {
				return math.MinInt32
			}
			if v > math.MaxInt32 {
				return math.MaxInt32
			}
			return v
		}
	}

	if m < 0 {
		return math.MinInt32
	} else {
		return math.MaxInt32
	}
}