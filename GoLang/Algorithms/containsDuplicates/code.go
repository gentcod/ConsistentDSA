package containsDuplicates

import (
	"fmt"
)

func Solve(arr []int) {
	fmt.Println(containsDuplicates(arr))
}

func containsDuplicates(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	unique := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		_, ok := unique[nums[i]]
		if ok {
			return true
		}
		unique[nums[i]] = true
	}
	return false
}