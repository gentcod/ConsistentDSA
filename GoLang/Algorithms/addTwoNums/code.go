package addtwonumbers

import (
	"fmt"

	"github.com/gentcod/ConsistentDSA/helper"
)

// Leetcode Problem: https://leetcode.com/problems/add-two-numbers/
func Solve(l1 *helper.ListNode, l2 *helper.ListNode) {
	node := addTwoNumbers(l1, l2)
	fmt.Println(helper.GetListNodeValues(node))
}

func addTwoNumbers(l1 *helper.ListNode, l2 *helper.ListNode) (res *helper.ListNode) {
	n1, n2, arr := l1, l2, []int{}
	n1Val, n2Val := 0, 0

	for n1 != nil || n2 != nil {
		if n1 != nil {
			n1Val = n1.Val
		}
		if n2 != nil {
			n2Val = n2.Val + n2Val
		} else if n1 == nil && n2Val == 1 {
			n2Val = 10
		}

		sum := n1Val + n2Val
		if sum >= 10 {
			sum = sum - 10
			n2Val = 1
		} else {
			n2Val = 0
		}
		arr = append(arr, sum)
		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}

		if n1 == nil && n2 == nil && n2Val > 0 {
			arr = append(arr, n2Val)
		}
		n1Val = 0
	}

	for n := len(arr) - 1; n >= 0; n-- {
		node := &helper.ListNode{
			Val:  arr[n],
			Next: nil,
		}
		node.Next = res
		res = node
	}

	return
}
