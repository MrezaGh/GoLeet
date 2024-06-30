package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	return dfs(node, visited)
}

func dfs(node *Node, visited map[*Node]*Node) *Node {
	if n, ok := visited[node]; ok {
		return n
	}

	newNode := &Node{Val: node.Val}
	visited[node] = newNode

	for i := 0; i < len(node.Neighbors); i++ {
		if n, ok := visited[node.Neighbors[i]]; ok {
			newNode.Neighbors = append(newNode.Neighbors, n)
			continue
		}
		newNode.Neighbors = append(newNode.Neighbors, dfs(node.Neighbors[i], visited))
	}

	return newNode
}

func main() {

}
