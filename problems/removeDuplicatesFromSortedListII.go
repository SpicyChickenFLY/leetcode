package problems

import "fmt"

func log(head, slow, fast *ListNode) {
    fmt.Print(strListWithPointers(head, []*ListNode{slow, fast}))
}

func deleteDuplicates(head *ListNode) *ListNode {
    result := &ListNode{0, head}

    slow := result
    // log(result, slow, head)

    for head != nil {

        // 快指针 验值 过重
        for (head.Next != nil && head.Next.Val == head.Val) {
            head = head.Next
        }
        // log(result, slow, head)

        // 快慢指针 判重
        if (slow.Next == head) {
            // 不重复 满指针推进, 快指针重置
            slow = head
            head = slow.Next
        } else {
            // 重复 ByPass 
            slow.Next = head.Next
            head = slow.Next
        }
        // log(result, slow, head)

    }
    return result.Next
}
