package main

import "fmt"

// TreeNode - Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateSubTree(left, right int) []*TreeNode {
	if left == right {
		return []*TreeNode{
			&TreeNode{left, nil, nil},
		}
	}
	finalResult := []*TreeNode{}
	for i := left; i <= right; i++ {
		var leftSubResult, rightSubResult []*TreeNode
		if i == left {
			leftSubResult = []*TreeNode{nil}
			rightSubResult = generateSubTree(i+1, right)
		} else if i == right {
			leftSubResult = generateSubTree(left, i-1)
			rightSubResult = []*TreeNode{nil}
		} else {
			leftSubResult = generateSubTree(left, i-1)
			rightSubResult = generateSubTree(i+1, right)
		}
		for _, lr := range leftSubResult {
			for _, rr := range rightSubResult {
				finalResult = append(finalResult, &TreeNode{i, lr, rr})
			}
		}
	}
	return finalResult
}

func generateTrees(n int) []*TreeNode {
	return generateSubTree(1, n)
}

func main() {
	for i := 1; i < 4; i++ {
		fmt.Println(i, generateTrees(i))
	}
	// generateTrees(2)
}
