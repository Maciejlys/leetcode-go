package leetcode

import "github.com/Maciejlys/leetcode-go/structures"

type TreeNode = structures.TreeNode

func hasPathSum(root *TreeNode, targetSum int) bool {

	var dfs func(node *TreeNode, sum int) bool
	dfs = func(node *TreeNode, sum int) bool {
		if node == nil {
			return false
		}
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			return sum == targetSum
		}
		return dfs(node.Left, sum) || dfs(node.Right, sum)
	}

	return dfs(root, 0)
}
