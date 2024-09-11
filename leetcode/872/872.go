package leetcode

import "github.com/Maciejlys/leetcode-go/structures"

type TreeNode = structures.TreeNode

func traverse(node *TreeNode, leaf chan int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		leaf <- node.Val
		return
	}

	if node.Left != nil {
		traverse(node.Left, leaf)
	}

	if node.Right != nil {
		traverse(node.Right, leaf)
	}
}

func compareLeaves(leaf1, leaf2 chan int, done chan bool) {
	for {

		l1, ok1 := <-leaf1
		l2, ok2 := <-leaf2

		if ok1 != ok2 {
			done <- false
			return
		}

		if !ok1 && !ok2 {
			done <- true
			return
		}

		if l1 != l2 {
			done <- false
			return
		}
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := make(chan int)
	leaf2 := make(chan int)
	done := make(chan bool)

	go func() {
		traverse(root1, leaf1)
		close(leaf1)
	}()

	go func() {
		traverse(root2, leaf2)
		close(leaf2)
	}()

	go compareLeaves(leaf1, leaf2, done)

	return <-done
}
