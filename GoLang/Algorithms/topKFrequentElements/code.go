package topkfrequent

import "fmt"

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

	for i := max; len(result) != k; i-- {
		for j := range freq[i] {
			if len(result) != k {
				result = append(result, freq[i][j])
			}
		}
	}

	return result
}


// ---------------- ROUGH WORK, passes only 16/21 test cases, due to time complexity
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