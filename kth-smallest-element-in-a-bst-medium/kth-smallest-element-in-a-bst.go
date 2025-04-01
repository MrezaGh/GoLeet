package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var dfs func(node *TreeNode, k int) (int, int)

	dfs = func(node *TreeNode, k int) (int, int) {
		if node == nil {
			return 0, -1
		}
		leftSize, leftItem := dfs(node.Left, k)
		rightSize, rightItem := dfs(node.Right, k-leftSize-1)
		fmt.Println("node:", node.Val, "left-size:", leftSize, "right-size:", rightSize)
		size := leftSize + rightSize + 1
		if size < k {
			return size, -1
		}

		if leftSize >= k {
			return size, leftItem
		} else if leftSize == k-1 {
			return size, node.Val
		} else {
			return size, rightItem
		}

	}

	_, item := dfs(root, k)

	return item
}
func main() {

}
