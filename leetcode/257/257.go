package leetcode

import (
	"github.com/Maciejlys/leetcode-go/structures"
	"strconv"
)

type TreeNode = structures.TreeNode

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)

	var dfs func(node *TreeNode, curr string)
	dfs = func(node *TreeNode, curr string) {
		if node == nil {
			return
		}
		curr += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			result = append(result, curr)
			return
		}
		curr += "->"
		dfs(node.Left, curr)
		dfs(node.Right, curr)
	}

	dfs(root, "")

	return result
}
