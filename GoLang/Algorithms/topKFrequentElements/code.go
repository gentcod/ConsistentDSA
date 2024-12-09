package topkfrequent

import "fmt"

// Leetcode Problem: https://leetcode.com/problems/top-k-frequent-elements/
func Solve(nums []int, k int) {
	fmt.Println(topKFrequentElements(nums, k))
}

func topKFrequentElements(nums []int, k int) (result []int) {
	if len(nums) == k {
		return nums
	}

	hashMap := make(map[int]int)
	max := 0
	for _, val := range nums {
		hashMap[val]++

		if hashMap[val] > max {
			max = hashMap[val]
		}
	}

	freq := make([][]int, max+1)
	for i, val := range hashMap {
		freq[val] = append(freq[val], i)
	}

	for len(result) != k {
		for i := range freq[max] {
			if len(result) != k {
				result = append(result, freq[max][i])
			}
		}
		max--
	}

	return result
}


// ---------------- ROUGH WORK, passes 16/21 test cases, due to time complexity
// func topKFrequentElements(nums []int, k int) (result []int) {
// 	if len(nums) == k {
// 		return nums
// 	}

// 	hashMap := make(map[int]int)
// 	unique := make(map[int]bool)

// 	for _, val := range nums {
// 		hashMap[val]++
// 	}

// 	if len(hashMap) == k {
// 		for i := range hashMap {
// 			result = append(result, i)
// 		}
// 	} else {
// 		for i := 0; len(unique) != k; i++ {
// 			_, ok := unique[getHighest(hashMap)]
// 			if !ok {
// 				unique[getHighest(hashMap)] = true
// 			}
// 			delete(hashMap, getHighest(hashMap))
// 		}
// 	}

// 	for i := range unique {
// 		result = append(result, i)
// 	}

// 	return result
// }

// func getHighest(hashMap map[int]int) (key int) {
// 	highest := 1
// 	for i, val := range hashMap {
// 		if val > highest {
// 			highest = val
// 			key = i
// 		}
// 	}

// 	return key
// }

// SAMPLE DATA
// nums := []int {-1,1,4,-4,3,5,4,-2,3,-1}
// nums := []int {5,3,1,1,1,3,5,73,1}
// nums := []int {1,1,1,2,2,3,3}