package oddevenlinkedlist

import (
	"fmt"

	"github.com/gentcod/ConsistentDSA/helper"
)

// Leetcode Problem: https://leetcode.com/problems/odd-even-linked-list/
func Solve(head *helper.ListNode) {
	fmt.Println(oddEvenList(head))
}

func oddEvenList(head *helper.ListNode) (res *helper.ListNode) {
	odd := []int{}
	even := []int{}
	comb := []int{}
	cur := head
	c := 1

	for cur != nil {
		if c%2 == 0 {
			even = append(even, cur.Val)
		} else {
			odd = append(odd, cur.Val)
		}
		cur = cur.Next
		c++
	}

	comb = append(comb, odd...)
	comb = append(comb, even...)

	for n := len(comb) - 1; n >= 0; n-- {
		node := &helper.ListNode{
			Val:  comb[n],
			Next: nil,
		}
		node.Next = res
		res = node
	}
	return
}
