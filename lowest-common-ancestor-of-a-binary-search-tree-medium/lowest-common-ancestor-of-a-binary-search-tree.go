package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	cache := make(map[[2]int]bool)

	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if p.Val == node.Val || q.Val == node.Val {
			return node
		}
		pLeft := in(node.Left, p, cache)
		qLeft := in(node.Left, q, cache)

		if pLeft != qLeft {
			return node
		} else if pLeft {
			return dfs(node.Left)
		} else {
			return dfs(node.Right)
		}
	}

	return dfs(root)
}

func in(node, target *TreeNode, cache map[[2]int]bool) bool {
	if node == nil {
		return false
	} else if node.Val == target.Val {
		return true
	} else if inside, exist := cache[[2]int{node.Val, target.Val}]; exist {
		return inside
	}

	inside := in(node.Left, target, cache) || in(node.Right, target, cache)
	cache[[2]int{node.Val, target.Val}] = inside
	return inside
}

func main() {

}
