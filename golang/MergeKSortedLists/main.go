package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(vals []int) *ListNode {
	result := &ListNode{0, nil}
	cur := result
	for _, val := range vals {
		cur.Next = &ListNode{val, nil}
		cur = cur.Next
	}
	return result.Next
}

func newLists(vals [][]int) []*ListNode {
	result := []*ListNode{}
	for _, val := range vals {
		result = append(result, newList(val))
	}
	return result
}

func mergeKLists(lists []*ListNode) *ListNode {
	var resultHead *ListNode = nil

	for _, list := range lists {
		if list == nil {
			continue
		}

		if resultHead == nil {
			resultHead = list
			continue
		}

		resultCur, listHead, listCur := resultHead, list, list

		// 如果存在比起始点小的的子序列,先把这段序列插进来
		if listCur.Val < resultCur.Val {
			listHead = nil
		}
		for listCur.Next != nil && listCur.Next.Val < resultCur.Val {
			listCur = listCur.Next
		}
		if listHead == nil {
			resultHead = list
			listHead = listCur.Next
			listCur.Next = resultCur
			listCur = listHead
		}

		// resultCur去重

		for listHead != nil {
			for resultCur.Next != nil && resultCur.Next.Val <= listHead.Val {
				resultCur = resultCur.Next
			}
			for listCur.Next != nil && resultCur.Next != nil && listCur.Next.Val < resultCur.Next.Val {
				listCur = listCur.Next
			}
			temp := listHead
			listHead = listCur.Next
			if resultCur.Next != nil {
				listCur.Next = resultCur.Next
			} else {
				listHead = nil
			}
			resultCur.Next = temp
			// 推进两个游标
			resultCur = listCur.Next
			listCur = listHead
		}
	}

	return resultHead
}

func main() {
	// lists := newLists([][]int{{1}, {0}})
	// lists := newLists([][]int{ {1, 4, 5}, {1, 3, 4}, {2, 6} })
	// lists := newLists([][]int{ {1, 2, 3}, {4, 5, 6, 7} })
	lists := newLists([][]int{ {1, 4, 5}, {0, 2} })
    mergeKLists(lists)
}
