package utils

import "errors"

// Stack is a data structure interface
type Stack interface {
	IsEmpty() bool
	Push(value interface{})
	Top() (interface{}, error)
	Pop() (interface{}, error)
}

// MyStack my inplemention for Stack
type MyStack struct {
	top  *MyListNode
	size int
}

// NewMyStack initialize MyStack here.
func NewMyStack(initVals []interface{}) *MyStack {
	newStack := &MyStack{
		top: nil,
	}
	for _, initVal := range initVals {
		newStack.Push(initVal)
	}
	return newStack
}

// IsEmpty chech the top ptr
func (ms *MyStack) IsEmpty() bool {
	return ms.top == nil
}

// Push add new head as top ptr
func (ms *MyStack) Push(value interface{}) {
	ms.top = &MyListNode{Value: value, Next: ms.top}
}

// Top return the value of top ptr
func (ms *MyStack) Top() (interface{}, error) {
	if ms.top == nil {
		return nil, errors.New("Empty Stack")
	}
	return ms.top.Value, nil
}

// Pop move the top ptr to next an return current
func (ms *MyStack) Pop() (interface{}, error) {
	if ms.top == nil {
		return nil, errors.New("Empty Stack")
	}
	value := ms.top.Value
	ms.top = ms.top.Next
	return value, nil
}
