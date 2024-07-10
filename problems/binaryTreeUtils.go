package problems

// ListNode - Definition for singly-linked list.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func newTree(vals []*TreeNode) *TreeNode {
	for index, val := range vals {
		if val == nil || index == 0 {
			continue
		}
		quotient := (index - 1) / 2
		remainder := (index - 1) % 2
		if remainder == 0 {
			vals[quotient].Left = val
		} else {
			vals[quotient].Right = val
		}
	}
	return vals[0]
}

func compareTwoTree(t1, t2 *TreeNode) bool {
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	if !compareTwoTree(t1.Left, t2.Left) {
		return false
	}
	if !compareTwoTree(t1.Right, t2.Right) {
		return false
	}
	return true
}

//
// func strTree(head *ListNode) string {
//     result := ""
//     for head != nil {
//         result += fmt.Sprintf("->%d", head.Val)
//         head = head.Next
//     }
//     return result
// }
//
// func strListWithPointers(head *ListNode, pointers []*ListNode) string {
//     result := ""
//     for head != nil {
//         result += fmt.Sprintf("->%d", head.Val)
//         for i, p := range pointers {
//             if (p == head) {
//                 result += fmt.Sprintf("[%d]", i)
//             }
//         }
//         head = head.Next
//     }
//     return result
// }

// NOTE: 仅用两根指针进行二叉树遍历(root, temp)
func morrisTraversal(root *TreeNode) {
	var temp *TreeNode = nil
	for root != nil {
		if root.Left != nil {
			// NOTE: connect threading for root
			temp = root.Left
			for temp.Right != nil && temp.Right != root {
				temp = temp.Right
			}
			// NOTE: the threading already exists
			if temp.Right != nil {
				temp.Right = nil
				// System.out.println(root.val);
				root = root.Right
			} else {
				// NOTE: construct the threading
				temp.Right = root
				root = root.Left
			}
		} else {
			// System.out.println(root.val);
			root = root.Right
		}
	}
}
