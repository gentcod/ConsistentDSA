package encodedecode

import (
	"fmt"
	"strconv"
	// "strings"
)

// Leetcode Problem: https://leetcode.com/problems/encode-and-decode-strings/
func Solve(value []string) {
	fmt.Println(encode(value))
	fmt.Println(decode(encode(value)))
}

// TIME COMPLEXITY - 0(n), SPACE COMPLEXITY - 0(N)

func encode(value []string) (res string) {
	for _,s := range value {
		res += strconv.Itoa(len(s)) + "#" + s
	}
	
	return
}

func decode(value string) (res []string) {
	i := 0

	for i < len(value) {
		j := i
		for string(value[j]) != "#" {
			j++
		}
		l,_ := strconv.Atoi(string(value[i:j]))
		i = j + 1
		j = i + l
		res = append(res, value[i:j])
		i = j
	}

	return
}

// TIME COMPLEXITY - 0(n), SPACE COMPLEXITY - 0(N)

// func encode(value []string) (res string) {
// 	sizes := []int{}

// 	for _,s := range value {
// 		sizes = append(sizes, len(s))
// 	}
	
// 	for _,sz := range sizes {
// 		res += strconv.Itoa(sz) + "_"
// 	}

// 	res += "#"

// 	for _,s := range value {
// 		res += s
// 	}
	
// 	return
// }

// func decode(value string) (res []string) {
// 	arr := strings.Split(value, "_")
// 	chars := arr[len(arr) - 1][1:]
// 	keys := arr[:len(arr) - 1]

// 	for _,k := range keys {
// 		j, _ := strconv.Atoi(k)
// 		res = append(res, chars[:j])
// 		chars = chars[j:]
// 	}

// 	fmt.Println(arr, chars, keys)

// 	return
// }

// TODO: resolve test cases for : ["1,23","45,6","7,8,9"] 
// and ["","   ","!@#$%^&*()_+","LongStringWithNoSpaces","Another, String With, Commas"]