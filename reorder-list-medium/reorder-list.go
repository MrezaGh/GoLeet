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

func Reverse(node ListNode) (*ListNode, int) {
	length := 0
	originalHead := &node
	var pre *ListNode
	var cur *ListNode
	for originalHead != nil {
		cur = &ListNode{Val: originalHead.Val}
		cur.Next = pre
		pre = cur
		originalHead = originalHead.Next
		length += 1
	}
	return cur, length
}

func reorderList(head *ListNode) {
	reversedHead, length := Reverse(*head)
	normal := head
	newList := normal
	normal = normal.Next
	updated := newList
	step := 1
	for length > 1 {
		fmt.Println("period:", updated, "rev:", reversedHead, "normal:", normal)

		if step%2 == 1 {
			newList.Next = reversedHead
			reversedHead = reversedHead.Next
		} else {
			newList.Next = normal
			normal = normal.Next
		}

		newList = newList.Next
		//newList.Next = nil
		step += 1
		length -= 1
	}
	newList.Next = nil
	head = updated
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
	reversedHead, length := Reverse(*list)
	fmt.Println(reversedHead)
	fmt.Println(length)

	reorderList(list)
	fmt.Println("====")
	fmt.Println(list)
}
