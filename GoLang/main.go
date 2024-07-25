package main

import (
	"fmt"
	// "github.com/gentcod/ConsistentDSA/Algorithms/sorting"
	// "github.com/gentcod/ConsistentDSA/Algorithms/linkedList"
	// "github.com/gentcod/ConsistentDSA/DataStructures/linkedList"
	"github.com/gentcod/ConsistentDSA/Algorithms/removeDuplicates"
)

func main() {
	fmt.Println("Loading Programs......")
	// numbers := []int {2, 6, 1, 8, 10, 4, 1}
	nums := []int {1, 2, 1, 2, 3, 2}

	//Run your imported modules here - Check ref.txt
	// singlyLinkedList := linkedList.New(numbers)
	// fmt.Println(singlyLinkedList.Size())
	// fmt.Println(singlyLinkedList.IsEmpty())
	removeDuplicates.Solve(nums)
}