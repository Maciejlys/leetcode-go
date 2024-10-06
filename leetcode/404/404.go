package leetcode

import "github.com/Maciejlys/leetcode-go/structures"

type TreeNode = structures.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			sum += node.Left.Val
		}

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return sum
}
