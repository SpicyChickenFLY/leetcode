package main

import (
	"errors"
	"fmt"
	"reflect"
)

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

// =================================================

type MyQueue struct {
	stack1, stack2 Stack
	size           int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack1: NewMyStack(),
		stack2: NewMyStack(),
	}
}

/** Push element x to the back of queue. */
func (mq *MyQueue) Push(x int) {
	for !mq.stack2.IsEmpty() {
		temp, _ := mq.stack2.Pop()
		mq.stack1.Push(temp)
	}
	mq.stack1.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (mq *MyQueue) Pop() int {
	for !mq.stack1.IsEmpty() {
		temp, _ := mq.stack1.Pop()
		mq.stack2.Push(temp)
	}
	value, _ := mq.stack2.Pop()
	return int(reflect.ValueOf(value).Int())
}

/** Get the front element. */
func (mq *MyQueue) Peek() int {
	for !mq.stack1.IsEmpty() {
		temp, _ := mq.stack1.Pop()
		mq.stack2.Push(temp)
	}
	value, _ := mq.stack2.Top()
	return int(reflect.ValueOf(value).Int())
}

/** Returns whether the queue is empty. */
func (mq *MyQueue) Empty() bool {
	return mq.stack1.IsEmpty() && mq.stack2.IsEmpty()
}

func main() {
	stack := NewMyStack()
	stack.Push(1)
	fmt.Println(stack.Top())
	stack.Push(2)
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())

	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Peek())
	fmt.Println(queue.Empty())
}
