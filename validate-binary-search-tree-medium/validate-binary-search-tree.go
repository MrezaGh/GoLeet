package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, lowerBound, UpperBound int) bool
	dfs = func(node *TreeNode, lowerBound, UpperBound int) bool {
		if node == nil {
			return true
		}
		if (node.Left != nil && node.Left.Val >= node.Val) || (node.Right != nil && node.Right.Val <= node.Val) {
			return false
		}
		if node.Val >= UpperBound || node.Val <= lowerBound {
			return false
		}
		//newLB, newUB := max(node.Val, lowerBound), min(node.Val, UpperBound)
		return dfs(node.Left, lowerBound, node.Val) && dfs(node.Right, node.Val, UpperBound)
	}

	return dfs(root, -(2<<31 + 1), 2<<31)

}

func main() {
	//if math.MaxInt32 > 2<<31-1 {}
	fmt.Println(1 << 2)
	fmt.Println(math.MaxInt32, 2<<30)
	fmt.Println(-math.MaxInt32, -2<<31)
}
