package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node ListNode) String() string {
	head := node
	str := ""
	for head.Next != nil {
		str += strconv.Itoa(head.Val) + " "
		head = *head.Next
	}
	str += strconv.Itoa(head.Val)
	return str
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	iter := head
	for iter != nil {
		length += 1
		iter = iter.Next
	}
	delIndex := length - n
	fmt.Println("should delete index:", delIndex, "length:", length)
	var pre *ListNode
	cur := head
	for delIndex > 0 {
		cur, pre = cur.Next, cur
		delIndex -= 1
	}

	if pre == nil {
		head = cur.Next
	} else {
		pre.Next = cur.Next
	}

	return head
}

func main() {
	head := &ListNode{Val: -1}
	list := head
	for i := 0; i < 5; i++ {
		head.Next = &ListNode{Val: i}
		head = head.Next
	}
	list = list.Next

	fmt.Println(list)
	newHead := removeNthFromEnd(list, 5)
	fmt.Println(newHead)

}
