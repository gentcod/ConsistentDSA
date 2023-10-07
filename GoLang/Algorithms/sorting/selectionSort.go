package selectionSort

import "fmt"

func Solve() {
	numbers := []int {2, 6, 1, 8, 10, 4, 1}
	fmt.Println(selectionSort(numbers))
}

func minIndex(arr []int) int {
	min := 0
	for i:= 0; i < len(arr); i++ {
		if arr[i] < arr[min] {
			min = i
		}
	}

	return min
}

func selectionSort(arr []int) []int {
	sortedArray := []int {}
	arrCopy := arr

	for i:= 0; i < len(arr); i++ {
		sortedArray = append(sortedArray, arrCopy[minIndex(arrCopy)])
		arrCopy = append(arrCopy[:minIndex(arrCopy)], arrCopy[minIndex(arrCopy)+1:]...)
	}

	return sortedArray
}