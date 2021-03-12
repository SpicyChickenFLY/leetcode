package main

// TreeNode - Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursive(root, p, q *TreeNode) (*TreeNode, int) {
	// post order
	if root == nil {
		return nil, 0
	} else if root == p || root == q {
		if _, leftN := recursive(root.Left, p, q); leftN > 0 {
			return root, 2
		}
		if _, rightN := recursive(root.Right, p, q); rightN > 0 {
			return root, 2
		}
		return root, 1
	} else {
		leftP, leftN := recursive(root.Left, p, q)
		rightP, rightN := recursive(root.Right, p, q)
		if leftN == 1 && rightN == 1 {
			return root, 2
		}
		if leftN > 0 {
			return leftP, leftN
		}
		if rightN > 0 {
			return rightP, rightN
		}
		return nil, 0
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	result, _ := recursive(root, p, q)
	return result
}

func main() {

}
