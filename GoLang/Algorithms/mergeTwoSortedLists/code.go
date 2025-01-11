package mergetwosortedlists

import (
	"fmt"

	"github.com/gentcod/ConsistentDSA/helper"
)

// Leetcode Problem: https://leetcode.com/problems/merge-two-sorted-lists/
func Solve(list1 *helper.ListNode, list2 *helper.ListNode) {
	node := mergeTwoLists(list1, list2)
	fmt.Println(helper.GetListNodeValues(node))
}

func mergeTwoLists(list1 *helper.ListNode, list2 *helper.ListNode) (res *helper.ListNode) {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
