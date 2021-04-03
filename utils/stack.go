package utils

import "errors"

type Stack interface {
	IsEmpty() bool
	Push(value interface{})
	Top() (interface{}, error)
	Pop() (interface{}, error)
}

type MyStack struct {
	top  *ListNode
	size int
}

/** Initialize your data structure here. */
func NewMyStack(initVals []interface{}) *MyStack {
	newStack := &MyStack{
		top: nil,
	}
	for _, initVal := range InitVals {
		newStack.Push(initVal)
	}
	return newStack
}

func (ms *MyStack) IsEmpty() bool {
	return ms.top == nil
}

func (ms *MyStack) Push(value interface{}) {
	ms.top = &ListNode{Value: value, Next: ms.top}
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
