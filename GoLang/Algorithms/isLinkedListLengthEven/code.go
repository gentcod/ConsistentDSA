package islinkedlistlengtheven

import (
	"fmt"

	"github.com/gentcod/ConsistentDSA/helper"
)

func Solve(head *helper.ListNode) {
	fmt.Println(isLengthEven(head))
}

func isLengthEven(head *helper.ListNode) bool {
	cur := head
	c := 0

	for cur != nil {
		c++
		cur = cur.Next
	}
	return c % 2 == 0
}