package palindromelinkedlist

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/palindrome-linked-list/
func Solve(head *ListNode) {
	fmt.Println(isPalindrome(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	nums := []int{}
	nums = append(nums, head.Val)
	next := head.Next

	for next != nil {
		nums = append(nums, next.Val)
		nxt := next.Next
		next = nxt
	}

	i, j := 0, len(nums)-1

	for i < j {
		if nums[i] != nums[j] {
			return false
		}
		i = i + 1
		j = j - 1
	}

	return true
}
