package slidingwindow

import "fmt"

// Leetcode Problem: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func Solve(arr []int) {
	fmt.Println(maxProfit(arr))
}

func max(num1 int, num2 int) int {
	if num1 >= num2 {
		return num1
	} else {
		return num2
	}
}

func maxProfit(prices []int) int {
	l, r := 0, 1
	maxP := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxP = max(maxP, profit)
		} else {
			l = r
		}
		r += 1
	}

	return maxP
}