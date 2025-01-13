package main

import "fmt"

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(nums []int) *ListNode {
	var head *ListNode = &ListNode{
		Val:  0,
		Next: nil,
	}
	temp := head
	for _, num := range nums {
		node := &ListNode{
			Val:  num,
			Next: nil,
		}
		temp.Next = node
		temp = temp.Next
	}
	return head.Next
}

func printList(head *ListNode) {
	temp := head
	for temp != nil {
		fmt.Printf("%d->", temp.Val)
		temp = temp.Next
	}
	fmt.Println("NULL")
}

func deleteDuplicates(head *ListNode) *ListNode {
	temp := head
	for temp != nil && temp.Next != nil {
		for temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
			if temp.Next == nil {
				return head
			}
		}
		temp = temp.Next
	}
	return head
}

func main() {
	testCases := [][]int{
		{1, 1, 2},
		{1, 1, 2, 3, 3},
		{1, 1, 1},
	}

	for i, testCase := range testCases {
		realTestCase := newList(testCase)
		fmt.Printf("Index:%d\n", i)
		printList(realTestCase)
		deleteDuplicates(realTestCase)
		printList(realTestCase)
	}

}
