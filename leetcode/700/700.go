package leetcode

import "github.com/Maciejlys/leetcode-go/structures"

type TreeNode = structures.TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}

	return nil
}
