package sorting

import "fmt"

func MergeSort(arr []int) {
	fmt.Println(mergeSort(arr))
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid_index := len(arr) / 2
	left_arr := mergeSort(arr[:mid_index])
	right_arr := mergeSort(arr[mid_index:])
	sorted := []int {}

	left_index := 0
	right_index := 0

	for left_index < len(left_arr) && right_index < len(right_arr) {
		if left_arr[left_index] < right_arr[right_index] {
			sorted = append(sorted, left_arr[left_index])
			left_index += 1
		} else {
			sorted = append(sorted, right_arr[right_index])
			right_index += 1
		}
	}

	sorted = append(sorted, left_arr[left_index:]...)
	sorted = append(sorted, right_arr[right_index:]...)

	return sorted
}