package problems

import "fmt"

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(vals []int) *ListNode {
	result := &ListNode{0, nil}
	cur := result
	for _, val := range vals {
		cur.Next = &ListNode{val, nil}
		cur = cur.Next
	}
	return result.Next
}

func compareTwoList(l1, l2 *ListNode) bool {
    for l1 != nil || l2 != nil {
        if l1 == nil || l2 == nil {
            return false
        }
        if l1.Val != l2.Val {
            return false
        }
        l1 = l1.Next
        l2 = l2.Next
    }
    return true
}

func strList(head *ListNode) string {
    result := ""
    for head != nil {
        result += fmt.Sprintf("->%d", head.Val)
        head = head.Next
    }
    return result
}

func strListWithPointers(head *ListNode, pointers []*ListNode) string {
    result := ""
    for head != nil {
        result += fmt.Sprintf("->%d", head.Val)
        for i, p := range pointers {
            if (p == head) {
                result += fmt.Sprintf("[%d]", i)
            }
        }
        head = head.Next
    }
    return result
}
