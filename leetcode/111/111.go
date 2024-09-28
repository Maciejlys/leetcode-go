package leetcode

import "github.com/Maciejlys/leetcode-go/structures"

type TreeNode = structures.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{}
	queue = append(queue, root)

	res := 0

	for len(queue) > 0 {
		queueSize := len(queue)

		for i := 0; i < queueSize; i++ {
			dequeuedItem := queue[0]
			queue = queue[1:]

			if i == 0 {
				res++
			}

			if dequeuedItem.Left == nil && dequeuedItem.Right == nil {
				return res
			}

			if dequeuedItem.Left != nil {
				queue = append(queue, dequeuedItem.Left)
			}

			if dequeuedItem.Right != nil {
				queue = append(queue, dequeuedItem.Right)
			}
		}
	}

	return res
}
