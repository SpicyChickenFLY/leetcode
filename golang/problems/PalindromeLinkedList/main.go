package main

import (
	"errors"
	"fmt"
)

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack interface {
	IsEmpty() bool
	Push(value interface{})
	Top() (interface{}, error)
	Pop() (interface{}, error)
}

type LinkListNode struct {
	Value interface{}
	Next  *LinkListNode
}

type MyStack struct {
	top  *LinkListNode
	size int
}

/** Initialize your data structure here. */
func NewMyStack() *MyStack {
	return &MyStack{
		top: nil,
	}
}

func (ms *MyStack) IsEmpty() bool {
	return ms.top == nil
}

func (ms *MyStack) Push(value interface{}) {
	if ms.top == nil {
		ms.top = &LinkListNode{Value: value}
	} else {
		ms.top = &LinkListNode{Value: value, Next: ms.top}
	}
}

func (ms *MyStack) Top() (interface{}, error) {
	if ms.top == nil {
		return nil, errors.New("Empty Stack")
	}
	return ms.top.Value, nil
}

func (ms *MyStack) Pop() (interface{}, error) {
	if ms.top == nil {
		return nil, errors.New("Empty Stack")
	}
	value := ms.top.Value
	ms.top = ms.top.Next
	return value, nil
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slowPtr, fastPtr := head, head
	stack := NewMyStack()
	// fastPtr推进速度为slowPtr的一倍
	for fastPtr != nil {
		fastPtr = fastPtr.Next
		// slowPtr节点的值压栈
		stack.Push(slowPtr.Val)
		slowPtr = slowPtr.Next
		if fastPtr != nil {
			fastPtr = fastPtr.Next
		} else {
			stack.Pop()
		}
	}

	// 此时slowPtr在列表的一半位置处
	for slowPtr != nil {
		// 匹配slowPtr节点值与栈顶元素
		val, _ := stack.Pop()
		if slowPtr.Val != val.(int) {
			// 如果匹配则出栈继续
			return false
		}
		slowPtr = slowPtr.Next
	}
	return slowPtr == nil
}

func main() {
	testCases := []*ListNode{
		{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}},
		{1, &ListNode{2, nil}},
		{1, &ListNode{0, &ListNode{0, nil}}},
	}

	for _, testCase := range testCases {
		fmt.Println(isPalindrome(testCase))
	}
}
