package helper

type ListNode struct {
	Val  int
	Next *ListNode
}

// InitListNode helps to create a LinkedList
func InitListNode(nums []int) (res *ListNode) {
	for i := len(nums) - 1; i >= 0; i-- {
		node := &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		node.Next = res
		res = node
	}

	return
}

// GetListNodeValues helps to get the values of the nodes in a linkedlit
func GetListNodeValues(node *ListNode) (res []int) {
	cur := node
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return
}