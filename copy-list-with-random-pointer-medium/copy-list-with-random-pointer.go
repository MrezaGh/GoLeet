package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nodeStoreMap := make(map[*Node]*Node, 0)
	newHead := &Node{}
	pre := newHead
	oldHead := head
	for oldHead != nil {
		newNode := &Node{Val: oldHead.Val}
		newNode.Random = oldHead.Random
		pre.Next = newNode
		if _, ok := nodeStoreMap[oldHead]; !ok {
			nodeStoreMap[oldHead] = newNode
		}
		oldHead = oldHead.Next
		pre = pre.Next
	}
	//fmt.Println(nodeStoreMap)
	newHeadCursor := newHead.Next
	for newHeadCursor != nil {
		if node, ok := nodeStoreMap[newHeadCursor.Random]; ok {
			newHeadCursor.Random = node
		} else {
			newHeadCursor.Random = nil
		}
	}

	return newHead.Next
}
func main() {

}
