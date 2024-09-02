package leetcode

import "github.com/Maciejlys/leetcode-go/structures"

type TreeNode = structures.TreeNode

func dfs(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= min.Val {
		return false
	}

	if max != nil && root.Val >= max.Val {
		return false
	}

	return dfs(root.Left, min, root) && dfs(root.Right, root, max)
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, nil, nil)
}
