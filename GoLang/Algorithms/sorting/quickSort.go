package sorting

import "fmt"

func QuickSort(arr []int) {
	fmt.Println(quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) <= 1{
		return arr
	}

	pivot := arr[0]
	arrLesser := []int {}
	arrGreater := []int {}

	for i := 1; i < len(arr); i++ {
		if arr[i] <= pivot {
			arrLesser = append(arrLesser, arr[i])
		} else {
			arrGreater = append(arrGreater, arr[i])
		}
	}

	return append(append(quickSort(arrLesser), pivot), quickSort(arrGreater)...)
}