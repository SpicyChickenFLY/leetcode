package main

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	dict := make(map[*ListNode]bool)
	temp := headA
	for temp != nil {
		dict[temp] = true
		temp = temp.Next
	}
	temp = headB
	for temp != nil {
		if _, ok := dict[temp]; ok {
			return temp
		}
		temp = temp.Next
	}
	return nil
}
