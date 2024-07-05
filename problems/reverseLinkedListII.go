package problems

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil || head.Next == nil || left == right {
        return head;
    }
    var temp = head
    var l *ListNode
    var cnt int
    for cnt = 2; cnt < left; cnt++ {
        temp = temp.Next
    }
    if (left > 1) {
        l = temp
    }
    r := temp
    if l != nil {
        r = temp.Next
    } else {
        cnt = 1
    }
    pre := r
    cur := r.Next
    pre.Next = nil;
    for {
        temp = cur.Next;
        cur.Next = pre;
        pre = cur;
        cur = temp;
        cnt++
        if (cnt >= right) {
            r.Next = cur
            break
        }
    }
    if l != nil {
        l.Next = pre
    } else {
        head = pre
    }
    return head
}

