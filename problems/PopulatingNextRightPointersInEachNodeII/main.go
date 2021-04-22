package main

// Node - Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect1(root *Node) *Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	if root.Right != nil {
		temp := root
		for temp.Next != nil {
			if temp.Next.Left != nil {
				root.Right.Next = temp.Next.Left
				break
			} else if temp.Next.Right != nil {
				root.Right.Next = temp.Next.Right
				break
			}
			temp = temp.Next
		}
		connect(root.Right)
	}
	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			temp := root
			for temp.Next != nil {
				if temp.Next.Left != nil {
					root.Left.Next = temp.Next.Left
					break
				} else if temp.Next.Right != nil {
					root.Left.Next = temp.Next.Right
					break
				}
				temp = temp.Next
			}
		}
		connect(root.Left)
	}
	return root
}

func subConnect(root *Node, depth int, dict map[int]*Node) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	if root.Right != nil {
		root.Right.Next = dict[depth]
		dict[depth] = root.Right
		if _, ok := dict[depth+1]; !ok {
			dict[depth+1] = nil
		}
		subConnect(root.Right, depth+1, dict)
	}
	if root.Left != nil {
		root.Left.Next = dict[depth]
		dict[depth] = root.Left
		if _, ok := dict[depth+1]; !ok {
			dict[depth+1] = nil
		}
		subConnect(root.Left, depth+1, dict)
	}
}

func connect(root *Node) *Node {
	dict := make(map[int]*Node)
	dict[1] = nil
	subConnect(root, 1, dict)
	return root
}

func main() {
	testCases := []*Node{
		&Node{1, &Node{2, &Node{4, nil, nil, nil}, &Node{5, nil, nil, nil}, nil}, &Node{3, nil, &Node{7, nil, nil, nil}, nil}, nil},
	}
	for _, testCase := range testCases {
		connect(testCase)
	}

}
