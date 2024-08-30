package leetcode

import "github.com/Maciejlys/leetcode-go/structures"

type TreeNode = structures.TreeNode

func findDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	left := findDepth(root.Left, depth+1)
	right := findDepth(root.Right, depth+1)

	if left > right {
		return left
	} else {
		return right
	}
}

func maxDepth(root *TreeNode) int {
	return findDepth(root, 0)
}
