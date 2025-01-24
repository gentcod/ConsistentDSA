package validatebst

import (
	"fmt"

	"github.com/gentcod/ConsistentDSA/helper"
)

// Leetcode Problem: https://leetcode.com/problems/validate-binary-search-tree/
func Solve(root *helper.TreeNode) {
	fmt.Println(isValidBST(root))
}

func isValidBST(root *helper.TreeNode) bool {
	return comp(root, nil, nil)
}

func comp(n, min, max *helper.TreeNode) bool {
	if n == nil {
		return true
	}
	if min != nil && n.Val <= min.Val {
		return false
	}
	if max != nil && n.Val >= max.Val {
		return false
	}
	return comp(n.Left, min, n) && comp(n.Right, n, max)
}
