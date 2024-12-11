package oddevenlinkedlist

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/odd-even-linked-list/
func Solve(head *ListNode) {
	fmt.Println(oddEvenList(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) (res *ListNode) {
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
		node := &ListNode{
			Val:  comb[n],
			Next: nil,
		}
		node.Next = res
		res = node
	}
	return
}
