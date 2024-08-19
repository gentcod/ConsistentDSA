package topkfrequent

import "fmt"

func Solve(nums []int, k int) {
	fmt.Println(topKFrequentElements(nums, k))
}

func topKFrequentElements(nums []int, k int) (result []int) {
	if len(nums) == k {
		return nums
	}

	// CREATE AN HASHMAP TO DETERMINE THE NUMBER OF TIMES EACH ELEMENT APPEAR
	hashMap := make(map[int]int)

	// INITIAL A MAX VARIABLE WHICH WILL BE THE BOUND FOR FREQUENCY ARRAY
	max := 0

	for _, val := range nums {
		hashMap[val]++

		if hashMap[val] > max {
			max = hashMap[val]
		}
	}

	// FREQUENCY USES THE INDEX OF EACH CHILD ARRAY/LIST AS REF TO ELEMENTS BASED ON NUMBER OF OCCURENCES
	// IF AN ELEMENT foo OCCURS 2 TIMES freq[2] -> [foo]
	freq := make([][]int, max+1)

	for i, val := range hashMap {
		freq[val] = append(freq[val], i)
	}

	for i := max; len(result) != k; i--{
		for j := range freq[i] {
			if len(result) == k {
				return result
			}
			result = append(result, freq[i][j])
		}
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