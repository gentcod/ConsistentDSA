package twoSum

import "fmt"

// Leetcode Problem: https://leetcode.com/problems/two-sum/
func Solve(arr []int, target int) {
	fmt.Print(twoSum(arr, target))
}

func twoSum(arr []int, target int) (res []int) {
	hashMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		diff := target - arr[i]
		num, ok := hashMap[arr[i]]

		if ok {
			res = append(res, num, i)
			return
		}

		hashMap[diff] = i
	}

	return
}