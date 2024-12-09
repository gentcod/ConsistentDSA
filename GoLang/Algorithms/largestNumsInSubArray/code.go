package largestNum

import (
	"fmt"
)

//PROBLEM:
/**
 * Return an array consisting of the largest number from each provided sub-array. For simplicity, the provided array will contain exactly 4 sub-arrays.

   Remember, you can iterate through an array with a simple for loop, and access each member with array syntax arr[i].
 */
//SOLUTION
func Solve(arr [][]int)  {
	fmt.Println(largestNumInSubArray(arr))
}

func largestNumInSubArray(arr [][]int) []int{
	max := []int{}

	for i := 0; i < len(arr); i++ {
		max = append(max, findMax(arr[i]))
	}
	return max
}

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// Test data: arr := [][]int{{4, 5, 1, 3}, {13, 27, 18, 26}, {32, 35, 37, 39}, {1000, 1001, 857, 1}}