package trappingwater

import (
	"fmt"
)

// PROBLEM lEETCODE - https://leetcode.com/problems/trapping-rain-water/description/

func Solve(height []int) {
	fmt.Println(trappingwater(height))
}

func trappingwater(height []int) (res int) {
	l, r := 0, len(height) - 1
	lMax, rMax := height[l], height[r]

	for l < r {
		if lMax < rMax {
			l++
			lMax = max(lMax, height[l])
			res += lMax - height[l]
		} else {
			r--
			rMax = max(rMax, height[r])
			res += rMax - height[r]
		}
	}

	return
}