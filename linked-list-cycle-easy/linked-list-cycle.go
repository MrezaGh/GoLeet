package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if _, ok := visited[cur]; ok {
			return true
		} else {
			visited[cur] = true
		}
		cur = cur.Next
	}
	return false
}

func main() {

}
