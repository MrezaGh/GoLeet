package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		cur, cur.Next, pre = cur.Next, pre, cur
	}
	return pre
}

func main() {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := head
	for i := 1; i <= 5; i++ {
		cur.Next = &ListNode{
			Val:  i,
			Next: nil,
		}
		cur = cur.Next
	}

	cur = reverseList(head)
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}

}
