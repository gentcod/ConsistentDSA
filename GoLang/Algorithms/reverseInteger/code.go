package reverseinteger

import (
	"fmt"
	"math"
	"strconv"
)

// Leetcode Problem: https://leetcode.com/problems/reverse-integer/
func Solve(value int) {
	fmt.Println(reverse(value))
}

func reverse(x int) (res int) {
	if invalidInt(x) {
		return 0
	}

	xs := strconv.Itoa(x)
	xLen := len(xs) - 1
	i, j, negative := 0, xLen, false

	if string(xs[0]) == "-" {
		xs = xs[1:]
		xLen--
		j--
		negative = true
	}

	for i <= xLen {
		v, _ := strconv.Atoi(string(xs[j]))
		res += v * pow(10, j)
		i++
		j--
	}

	if negative {
		res *= -1
	}

	if invalidInt(res) {
		return 0
	}

	return
}

func pow(base, exp int) int {
	if exp == 0 {
		return 1
	}

	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func invalidInt(val int) bool {
	if val < math.MinInt32 || val > math.MaxInt32 {
		return true
	}
	return false
}
