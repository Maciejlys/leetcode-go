package leetcode

import "slices"

func getLeafs(node *TreeNode, curr *[]int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		*curr = append(*curr, node.Val)
		return
	}

	if node.Left != nil {
		getLeafs(node.Left, curr)
	}

	if node.Right != nil {
		getLeafs(node.Right, curr)
	}
	return
}

func leafSimilar2(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1 := make([]int, 0)
	leafs2 := make([]int, 0)
	getLeafs(root1, &leafs1)
	getLeafs(root2, &leafs2)
	return slices.Equal(leafs1, leafs2)
}
