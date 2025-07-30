package mineatingspeed

import (
	"fmt"
	"math"
)

// Leetcode Problem: https://leetcode.com/problems/koko-eating-bananas/
func Solve(piles []int, h int) {
	fmt.Println(minEatingSpeed(piles, h))
}

func minEatingSpeed(piles []int, h int) (res int) {
	l, r := 1, getMax(piles)
	res = r

	for l <= r {
		k := (l + r) / 2
		totalTime := 0

		for _, p := range piles {
			totalTime += int(math.Ceil(float64(p) / float64(k)))
		}

		if totalTime <= h {
			res = k
			r = k - 1
		} else {
			l = k + 1
		}
	}

	return res
}

func getMax(arr []int) (max int) {
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return
}
