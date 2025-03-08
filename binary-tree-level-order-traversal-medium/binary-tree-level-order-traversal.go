package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type item struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	levels := make([][]int, 0)
	frontier := []*item{&item{
		node:  root,
		level: 0,
	}}
	for len(frontier) > 0 {
		head, level := frontier[0].node, frontier[0].level
		frontier = frontier[1:]
		if head == nil {
			continue
		}
		if len(levels) < level+1 {
			levels = append(levels, []int{})
		}
		levels[level] = append(levels[level], head.Val)
		frontier = append(frontier,
			&item{
				node:  head.Left,
				level: level + 1,
			},
			&item{
				node:  head.Right,
				level: level + 1,
			})
	}
	return levels
}
