package invertree

import (
	"fmt"

	"github.com/gentcod/ConsistentDSA/helper"
)

// Leetcode Problem: https://leetcode.com/problems/invert-binary-tree/
func Solve(root *helper.TreeNode) {
	node := invertTree(root)
	fmt.Println(node)
}

func invertTree(root *helper.TreeNode) *helper.TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		root.Left = invertTree(root.Left)
		root.Right = invertTree(root.Right)
	}
	return root
}
