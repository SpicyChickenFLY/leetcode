package utils

// List is a interface
type List interface {
	GetHead()
	GetTail()
	PushFront()
	PopFront()
	PushBack()
	PopBack()
}

// MyListNode is a node structure for linklisk
type MyListNode struct {
	Value interface{}
	Next  *MyListNode
}

// MyList is my inplemention for linkList
type MyList struct {
	Head *MyListNode
}

// NewMyList initalize new MyList
func NewMyList(initVals []interface{}) *MyList {
	var newHead *MyListNode
	for i := len(initVals) - 1; i >= 0; i-- {
		newHead = &MyListNode{Value: initVals[i], Next: newHead}
	}
	return &MyList{Head: newHead}
}

// Print all nodes in list
func (l *MyListNode) Print() {}

// 判断是否有回路
