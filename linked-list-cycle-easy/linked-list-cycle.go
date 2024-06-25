package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//func hasCycle(head *ListNode) bool {
//	visited := make(map[*ListNode]bool)
//	cur := head
//	for cur != nil {
//		if _, ok := visited[cur]; ok {
//			return true
//		} else {
//			visited[cur] = true
//		}
//		cur = cur.Next
//	}
//	return false
//}

func hasCycle(head *ListNode) bool {
	// constant space
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next

	for slow != fast {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {

}
