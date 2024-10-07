package leetcode

import (
	"math"

	"github.com/Maciejlys/leetcode-go/structures"
)

type TreeNode = structures.TreeNode

func dfs(node *TreeNode, depth int, sums map[int]int) {
	if node == nil {
		return
	}

	sums[depth] += node.Val

	dfs(node.Left, depth+1, sums)
	dfs(node.Right, depth+1, sums)
}

func maximum(sums map[int]int) (key int) {
	m := math.MinInt32

	for k, v := range sums {
		if v > m {
			key = k
			m = v
		}
	}

	return key
}

func maxLevelSum(root *TreeNode) int {
	sums := make(map[int]int)
	dfs(root, 1, sums)
	return maximum(sums)
}
