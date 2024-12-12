package islinkedlistlengtheven

import (
	"fmt"
)

func Solve(head *ListNode) {
	fmt.Println(isLengthEven(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isLengthEven(head *ListNode) bool {
	cur := head
	c := 0

	for cur != nil {
		c++
		cur = cur.Next
	}
	return c % 2 == 0
}