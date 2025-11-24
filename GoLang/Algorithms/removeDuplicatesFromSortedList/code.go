package removeDuplicatesFromSortedList

import (
	"fmt"

	"github.com/gentcod/ConsistentDSA/helper"
)

// Leetcode Problem:
func Solve(head *helper.ListNode) {
	fmt.Println(helper.GetListNodeValues(removeDuplicatesFromSortedList(head)))
}

func removeDuplicatesFromSortedList(head *helper.ListNode) *helper.ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}