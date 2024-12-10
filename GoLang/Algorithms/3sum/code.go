package threesum

import (
	"fmt"
	"sort"
	// "strings"
)

// Leetcode Problem: https://leetcode.com/problems/3sum/
func Solve(nums []int) {
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) (res [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target, left, right := -nums[i], i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
	return
}

// func threeSum(nums []int) (res [][]int) {
// 	hash := make(map[rune][]int)
// 	i, j, k := 0, 1, 2
// 	for k < len(nums) {
// 		if nums[i]+nums[j]+nums[k] == 0 {
// 			arr := []int{nums[i], nums[j], nums[k]}
// 			hash[arrayToKeyCustom(arr)] = arr
// 		}
// 		j++
// 		k++
// 	}
// 	for _, arr := range hash {
// 		res = append(res, arr)
// 	}
// 	return
// }

// func arrayToKeyCustom(arr []int) rune {
// 	var r rune
// 	for _, n := range arr {
// 		str := strconv.Itoa(n)
// 		r += rune(str[0])
// 	}

// 	return r
// }
