package binarysearch

import "fmt"

func SolveBinary(list []int, num int) {
	fmt.Println(binarySearch(list, num))
}

func SolveRecurssiveBinary(list []int, num int) {
	fmt.Println(recurssiveBinarySearch(list, num))
}

func binarySearch(list []int, num int) bool {
	first := 0
	last := len(list) - 1

	for first <= last {
		mid := (first + last) / 2

		fmt.Println("first: ", first, "mid: ", mid, "last: ", last)

		if list[mid] == num {
			return true
		} else if list[mid] < num {
			first = mid + 1
		} else if list[mid] > num {
			last = mid - 1
		}
	}

	return false
}

func recurssiveBinarySearch(list []int, num int) bool {
	first := 0
	last := len(list) - 1

	for first <= last {
		mid := (first + last) / 2
		
		if list[mid] == num {
			return true
		} else if list[mid] < num {
			return recurssiveBinarySearch(list[mid+1:], num)
		} else {
			return recurssiveBinarySearch(list[:mid], num)
		}
	}

	return false
}