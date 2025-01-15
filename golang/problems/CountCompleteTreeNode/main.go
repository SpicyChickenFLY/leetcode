package main

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return recursive(root, 1)
}

func recursive(root *TreeNode, val int) int {
	if root.Right == nil {
		if root.Left == nil {
			return val
		}
		return val * 2
	}
	leftVal := recursive(root.Left, val*2)
	rightVal := recursive(root.Right, val*2+1)
	if leftVal > rightVal {
		return leftVal
	}
	return rightVal
}

func main() {

}
