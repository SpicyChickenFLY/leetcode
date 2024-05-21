package main

import "fmt"

// TreeNode - Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subBTPaths(root *TreeNode, parentPath string) []string {
	if root == nil {
		return []string{}
	}
	if parentPath == "" {
		parentPath = fmt.Sprintf("%d", root.Val)
	} else {
		parentPath = fmt.Sprintf("%s->%d", parentPath, root.Val)
	}
	if root.Left == nil && root.Right == nil {
		return []string{parentPath}
	}
	result := subBTPaths(root.Left, parentPath)
	return append(
		result, subBTPaths(root.Right, parentPath)...)
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	result := subBTPaths(root, "")
	return result
}

func main() {
	testCases := []*TreeNode{
		{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}},
		{1, nil, nil},
	}
	for _, testCase := range testCases {
		fmt.Println(binaryTreePaths(testCase))
	}
}
