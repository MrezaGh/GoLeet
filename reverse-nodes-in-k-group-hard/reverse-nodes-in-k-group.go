package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func String(node *ListNode) string {
	str := ""
	head := node
	for head != nil {
		str += fmt.Sprintf("%d ->", head.Val)
		head = head.Next
	}
	return str
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	var ans *ListNode
	dummyHead := head
	var earlierLast *ListNode
	for {
		// fmt.Println(String(ans))
		first, last, next := reverseK(dummyHead, k)

		if first == nil {
			return ans
		}
		if ans == nil {
			ans = first
		}
		if earlierLast != nil {
			earlierLast.Next = first
		}
		last.Next = next
		dummyHead = next
		earlierLast = last
	}
}

func reverseK(head *ListNode, k int) (*ListNode, *ListNode, *ListNode) {
	next := head
	for k1 := k; k1 > 0; k1-- {
		if next == nil {
			return nil, nil, nil
		}
		next = next.Next
	}
	var pre *ListNode
	cur := head
	for k2 := k; k2 > 0; k2-- {
		cur, cur.Next, pre = cur.Next, pre, cur
	}
	// fmt.Println("rev:", String(pre))
	// fmt.Println("next:", String(next))
	return pre, head, next

}
