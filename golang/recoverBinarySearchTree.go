package golang

func recoverTree(root *TreeNode) {
    var curr = root
	var prev *TreeNode = nil
	var first *TreeNode = nil
	var secend *TreeNode = nil
	for curr != nil {
		if curr.Left != nil {

			// 寻找curr的前驱节点(左子树的最右节点)
            temp := curr.Left
			for temp.Right != nil && temp.Right != curr {
				temp = temp.Right
			}

			// 已经建立过curr与前驱节点的联系
			if temp.Right != nil {

				// fmt.Println(curr.val);
				if prev != nil && prev.Val > curr.Val {
					secend = curr
					if first == nil {
						first = prev
					}
				}
				prev = curr

				temp.Right = nil // 摧毁临时前驱指针
				curr = curr.Right

				// 尚未建立过前驱
			} else {
				temp.Right = curr // 创建临时前驱指针
				curr = curr.Left
			}

        // curr不存在左子树了,向右移动
		} else {

			// fmt.Println(curr.val);
			if prev != nil && prev.Val > curr.Val {
				secend = curr
				if first == nil {
					first = prev
				}
			}
			prev = curr

			curr = curr.Right // 向右移动
		}
	}

    // 交换两个乱序节点
    swap := first.Val
    first.Val = secend.Val
    secend.Val = swap
}
