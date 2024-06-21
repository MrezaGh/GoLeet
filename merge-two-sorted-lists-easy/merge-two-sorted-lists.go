package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{Val: -1}
	newList := head
	head1, head2 := list1, list2

	for head1 != nil || head2 != nil {
		if head1 != nil && head2 != nil {
			if head1.Val <= head2.Val {
				head.Next = head1
				head1 = head1.Next
			} else {
				head.Next = head2
				head2 = head2.Next
			}
		} else if head1 == nil {
			head.Next = head2
			head2 = head2.Next
		} else if head2 == nil {
			head.Next = head1
			head1 = head1.Next
		}
		head = head.Next
	}
	return newList.Next
}

func main() {

}
