import java.util.List;

public class ListNode {
    int val;
    ListNode next;

    ListNode(int val) {
        this.val = val;
    }
    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}

class Solution {
    // Runtime: 100.00%, Memory Usage: 98.92%
    public ListNode reverseList1(ListNode head) {
        if (head == null || head.next == null)
            return head;
        ListNode newHead = reverseList1(head.next);
        head.next.next = head;
        head.next = null;
        return newHead;
    }

    // Runtime: 100.00%, Memory Usage: 98.56%
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null)
            return head;
        ListNode pre = head, cur = head.next, temp = null;
        pre.next = null;
        while (true) {
            temp = cur.next;
            cur.next = pre;
            if (temp != null) {
                pre = cur;
                cur = temp;
            } else {
                return cur;
            }
        }
        
    }
}