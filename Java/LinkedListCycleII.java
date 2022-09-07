// Definition for singly-linked list.
class ListNode {
    int val;
    ListNode next;
    
    ListNode(int val) {
        this.val = val;
        this.next = null;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
};

class Solution {
    // Runtime: 20.20%, Memory Usage: 6.32%
    public ListNode detectCycle(ListNode head) {
        if (head == null || head.next == null || head.next.next == null) {
            return null;
        }
        ListNode fast_pointer = head.next.next;
        ListNode slow_pointer = head.next;
        boolean flag = false;
        int count = 0;
        while (fast_pointer.next != null && fast_pointer.next.next != null && slow_pointer.next != null) {
            if (flag)
                count++;
            System.out.println(count);
            if (fast_pointer == slow_pointer) {
                if (flag) {
                    ListNode start = head, end = start;
                    for (int i = 0; i < count; i++)
                        end = end.next;
                    while (start != end) {
                        start = start.next;
                        end = end.next;
                    }
                    return start;
                } else
                    flag = true;
            }
            fast_pointer = fast_pointer.next.next;
            slow_pointer = slow_pointer.next;
        }
        return null;
    }
}

public class LinkedListCycleII {
    public static void main(String[] args) {
        Solution s = new Solution();
        ListNode test_case1 = new ListNode(3, new ListNode(2, new ListNode(0, new ListNode(-4))));
        test_case1.next.next.next.next = test_case1.next.next;
        System.out.println("1:" + s.detectCycle(test_case1).val); 
        ListNode test_case2 = new ListNode(1, new ListNode(2));
        test_case2.next.next = test_case2;
        System.out.println("2:" + s.detectCycle(test_case2).val);
    }
}