package golang

func _reverseKGroup(head *ListNode, k, cnt int) (*ListNode, *ListNode) {
	if cnt == 0 {
		cnt = k // 重置Group的计数器
	}

	// 递归到链表最后一个节点
	if head == nil || head.Next == nil {
		if cnt == 1 { // 最后一个批次满足一个组的数量,需要反转
			return head, nil // 这个节点就是这组的新头
		} else {
			return nil, head // 没有新头, 告知调用者,别反转
		}
	}

	newHead, tail := _reverseKGroup(head.Next, k, cnt-1)
	if cnt == 1 { // 处理新的一组
		return head, tail // 这个节点就是这组的新头
	} else if newHead == nil { // 这组没处理完, 且被告知不用反转
		tail = head // 更新尾巴
		return nil, tail // 接着告知调用者别反转
	} else { // 这组没处理完, 但是需要反转
		head.Next.Next = head // 反转节点
		if cnt == k { // 这组即将完成
			head.Next = tail // 接上尾巴
			return newHead, newHead // 更新新头和尾巴
		} else {
			return newHead, tail // 头尾不变继续
		}
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	head, _ = _reverseKGroup(head, k, k)
	return head
}
