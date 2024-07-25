package removeDuplicates

import (
	"fmt"
	"sort"
)

func Solve(arr []int) {
	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
		}
	}
	return k
}