package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth(root)

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := 0, 0
		if node.Left != nil {
			left = node.Left.Val + 1
		}
		if node.Right != nil {
			right = node.Right.Val + 1
		}
		return max(dfs(node.Left), dfs(node.Right), left+right)
	}

	return dfs(root)

}

func depth(node *TreeNode) int {
	if node == nil {
		return -1
	}

	node.Val = max(depth(node.Left), depth(node.Right)) + 1
	return node.Val
}

func main() {

}
