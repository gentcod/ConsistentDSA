package twoSum

import "fmt"

// Leetcode Problem: https://leetcode.com/problems/two-sum/
func Solve(arr []int, target int) {
	fmt.Print(twoSum(arr, target))
}

func twoSum(arr []int, target int) []int {
	hashMap := make(map[int]int)
	indexArr := []int{}

	for i := 0; i < len(arr); i++ {
		diff := target - arr[i]
		num, ok := hashMap[arr[i]]

		if ok {
			indexArr = append(indexArr, num, i)

			//Ensure only first two indexes are recorded
			if len(indexArr) <= 2 {
				return indexArr
			}
		}

		hashMap[diff] = i
	}

	return indexArr
}