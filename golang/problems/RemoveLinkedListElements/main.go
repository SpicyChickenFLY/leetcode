package main

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{0, head}
	temp := newHead
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
			continue
		}
		temp = temp.Next
	}
	return newHead.Next
}

func main() {
	testCasesList := []*ListNode{
		{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}},
		{7, &ListNode{7, &ListNode{7, &ListNode{7, nil}}}},
	}
	testCaseVal := []int{
		6,
		7,
	}
	for i := 0; i < len(testCasesList); i++ {
		removeElements(testCasesList[i], testCaseVal[i])
	}
}
