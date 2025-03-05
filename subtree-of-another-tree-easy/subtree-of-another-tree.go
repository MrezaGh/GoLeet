package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if same(root, subRoot) {
		return true
	} else {
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}
}

func same(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}
	return same(left.Left, right.Left) && same(left.Right, right.Right) && left.Val == right.Val

}
