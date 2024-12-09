package productexceptself

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/product-of-array-except-self/
func Solve(value []int) {
	fmt.Println(productExceptSelf(value))
}

// USING PREFIX AND SUFFIX INTUITION: prefix is the product of numbers before the indexed element
// while postfix isthe product of numbers after the element - TIME: O(n), SPACE: 0(1)

func productExceptSelf(value []int) []int {
	res := make([]int, len(value))
	for i := range res {
		res[i] = 1
	}

	prefix := 1
	postfix := 1

	for i,v := range value {
		res[i] = prefix
		prefix *= v
	}

	for i := len(value) - 1; i >= 0; i-- {
		res[i] = postfix * res[i]
		postfix *= value[i]
	}
	
	return res
}

// INITIAL INTUITION - TIME: O(n), SPACE: 0(n)

// func productExceptSelf(value []int) (result []int) {
// 	for i := range value {
// 		except := copy(value)
// 		except = append(except[:i], except[i+1:]...)
// 		result = append(result, prod(except))
// 	}
// 	return
// }

// func copy(val []int) []int {
// 	newVal := make([]int, len(val))
// 	for i, v := range val {
// 		newVal[i] = v
// 	}
// 	return newVal
// }

// func prod(val []int) int {
// 	tot := 1
// 	for _,v := range val {
// 		tot *= v
// 	}
// 	return tot
// }