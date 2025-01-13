package main

import "fmt"

// TreeNode - Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	recursion(root, &result)
	return result

}

func recursion(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	recursion(root.Left, result)
	recursion(root.Right, result)

	*result = append(*result, root.Val)
	// fmt.Println(result)
}

func main() {
	testCases := []*TreeNode{
		nil,
		{1,
			nil,
			&TreeNode{2,
				&TreeNode{3,
					nil,
					nil,
				},
				nil,
			},
		},
		{1,
			&TreeNode{2,
				nil,
				nil,
			},
			nil,
		},
	}
	for _, testCase := range testCases {
		fmt.Println(postorderTraversal(testCase))
	}
}
