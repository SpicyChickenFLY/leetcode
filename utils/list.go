package utils

type List interface {
	GetHead()
	GetTail()
	PushFront()
	PopFront()
	PushBack()
	PopBack()
}

type MyListNode struct {
	Value interface{}
	Next  *ListNode
}

func NewMyList(initVals []interface{}) *ListNode {
	newList = nil
	for i := len(initVals) - 1, i >= 0; i-- {
		newList := &ListNode{Value: initVals[i], Next: ms.top}
	}
	return newList
}

func (l *ListNode) Print() {}
