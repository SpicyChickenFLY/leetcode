package problems

func _delNodes(root *TreeNode, delMap map[int]bool) (child *TreeNode ,result []*TreeNode) {
	if root == nil {
		return nil, nil
	}

	left, leftResult:= _delNodes(root.Left, delMap)
	right, rightResult:= _delNodes(root.Right, delMap)
	result = append(leftResult, rightResult...)

	if delMap[root.Val] {
		if left != nil {
			result = append(result, left)
		}
		if right != nil {
			result = append(result, right)
		}
		return nil, result
	}

	root.Left = left
	root.Right = right
	return root, result
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	delMap := make(map[int]bool)
	for _, del := range to_delete {
		delMap[del] = true
	}
	root, result := _delNodes(root, delMap)
	if root != nil {
		return append(result, root)
	}
	return result
}
