package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var pre *ListNode
	ans := &ListNode{}
	head := ans
	for l1 != nil || l2 != nil {
		singleSum := head.Val
		if l1 == nil {
			singleSum += l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			singleSum += l1.Val
			l1 = l1.Next
		} else {
			singleSum += l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}
		head.Val = singleSum % 10
		head.Next = &ListNode{Val: singleSum / 10}
		pre = head
		head = head.Next
	}
	if head.Val == 0 {
		pre.Next = nil
	}

	return ans
}

func main() {

}
