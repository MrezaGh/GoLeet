package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, maxSeen int) int
	dfs = func(node *TreeNode, maxSeen int) int {
		if node == nil {
			return 0
		}
		current := 0
		if node.Val >= maxSeen {
			maxSeen = node.Val
			current++
		}
		return dfs(node.Left, maxSeen) + dfs(node.Right, maxSeen) + current

	}

	return dfs(root, -math.MaxInt32)
}

func main() {

}
