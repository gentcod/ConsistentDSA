package maxarea

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/container-with-most-water/
func Solve(height []int) {
	fmt.Println(maxArea(height))
}

func maxArea(heights []int) (area int) {
	l1, l2 := 0, len(heights)-1

	for l1 < l2 {
		newArea := minH(heights[l1], heights[l2]) * (l2 - l1)
		if newArea > area {
			area = newArea
		}
		if heights[l1] <= heights[l2] {
			l1++
		} else {
			l2--
		}
	}
	return
}

// func maxArea(height []int) (area int) {
// 	l1, l2 := 0, len(height) - 1
// 	pivot := (len(height) - 1)/2

// 	if len(height) <= 2 {
// 		return height[minH(l1,l1)]
// 	}

// 	for h := range height {
// 		minH := minH(height[h], height[l1])
// 		newArea := minH * (h - l1)

// 		if height[h] > height[l1] && pivot >= h{
// 			l1 = h
// 			area = minH * (l2 - l1)
// 		}
	
// 		if newArea > area && pivot <= h {
// 			l2 = h
// 			area = minH * (l2 - l1)
// 		}
// 	}
// 	return
// }

func minH(l, r int) int {
	if l < r {
		return l
	} else {
		return r
	}
}
