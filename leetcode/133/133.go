package leetcode

import "github.com/Maciejlys/leetcode-go/structures"

type Node = structures.Node

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	visited := make(map[int]*Node)
	return dfs(node, visited)
}

func dfs(node *Node, visited map[int]*Node) *Node {
	i := node.Val
	if _, ok := visited[i]; ok {
		return visited[i]
	}

	clone := &Node{Val: i}
	visited[i] = clone
	for _, item := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfs(item, visited))
	}
	return clone
}
