package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	balanced, _ := depth(root)
	return balanced
}

func depth(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}
	lBalanced, lDepth := depth(node.Left)
	rBalanced, rDepth := depth(node.Right)
	diff := abs(lDepth - rDepth)

	if lBalanced && rBalanced && diff <= 1 {
		return true, 1 + max(lDepth, rDepth)
	} else {
		return false, 0
	}

}
func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func main() {

}
