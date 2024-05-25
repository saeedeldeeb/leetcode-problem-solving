package main

// ListNode is a struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var mergedHead = &ListNode{}
	currentMNode := mergedHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			currentMNode.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else if l1.Val == l2.Val {
			currentMNode.Next = &ListNode{Val: l1.Val}
			currentMNode.Next.Next = &ListNode{Val: l2.Val}
			currentMNode = currentMNode.Next
			l1 = l1.Next
			l2 = l2.Next
		} else {
			currentMNode.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		currentMNode = currentMNode.Next
	}
	if l1 != nil {
		currentMNode.Next = l1
	} else if l2 != nil {
		currentMNode.Next = l2
	}
	return mergedHead.Next
}
