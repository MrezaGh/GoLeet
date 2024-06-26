package main

import "container/heap"

type minHeap []*ListNode

func (h minHeap) Len() int { return len(h) }

func (h minHeap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return h[i].Val < h[j].Val
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	// this is fast
	h := &minHeap{}

	result := &ListNode{}
	resultHead := result
	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}

	for h.Len() > 0 {
		minNode := heap.Pop(h).(*ListNode)
		resultHead.Next = &ListNode{Val: minNode.Val}
		resultHead = resultHead.Next
		minNode = minNode.Next
		if minNode != nil {
			heap.Push(h, minNode)
		}

	}

	return result.Next

}

//func mergeKLists(lists []*ListNode) *ListNode {
//
////	this is slow
//	result := &ListNode{}
//	resultHead := result
//	for {
//		currentMin, currentMinIndex := math.MaxInt, 0
//		for i, head := range lists {
//			heap.Push()
//			if head != nil && head.Val < currentMin {
//				currentMin = head.Val
//				currentMinIndex = i
//			}
//		}
//		if currentMin == math.MaxInt {
//			break
//		}
//		minNode := lists[currentMinIndex]
//		lists[currentMinIndex] = minNode.Next
//		resultHead.Next = &ListNode{Val: minNode.Val}
//		resultHead = resultHead.Next
//	}
//	return result.Next
//
//}

func main() {

}
