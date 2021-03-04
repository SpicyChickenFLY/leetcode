package main

// Node - Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// This solution doesnt break original input
func copyRandomList(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)
	temp := head
	for temp != nil {
		nodeMap[temp] = &Node{temp.Val, nil, nil}
		temp = temp.Next
	}
	temp = head
	for temp != nil {
		nodeMap[temp].Next = nodeMap[temp.Next]
		nodeMap[temp].Random = nodeMap[temp.Random]
		temp = temp.Next
	}
	return nodeMap[head]
}

func main() {
	testCase1 := &Node{7, &Node{13, &Node{11, &Node{10, &Node{1, nil, nil}, nil}, nil}, nil}, nil}
	testCase1.Random = nil
	testCase1.Next.Random = testCase1
	testCase1.Next.Next.Random = testCase1.Next.Next.Next.Next
	testCase1.Next.Next.Next.Random = testCase1.Next.Next
	testCase1.Next.Next.Next.Next.Random = testCase1

	copyRandomList(testCase1)
}
