class ListNode {
    public int val;
    public ListNode next;

    public ListNode(int val) {
        this.val = val;
        this.next = null;
    }
    public ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}

class util {
    public static void pr(ListNode head) {
        while (head != null) {
            System.out.print(head.val + "->");
            head = head.next;
        }
        System.out.println("null");
    }
}

class Solution {
    // RunTime: 38.65%, MemoryUsage: 100.00%
    public ListNode insertionSortList(ListNode head) {
        ListNode result = null;
        while (head != null) {
            if (result == null) {
                result = head;
                head = head.next;
                result.next = null;
            } else if (head.val < result.val) {
                ListNode temp = head;
                head = head.next;
                temp.next = result;
                result = temp;
            } else {
                ListNode p = result;
                while (p.next != null && head.val >= p.next.val) {
                    p = p.next;
                }
                ListNode temp = head;
                head = head.next;
                temp.next = p.next;
                p.next = temp;
            }
        }
        return result;
    }
}

public class InsertionSortList {
    public static void main(String[] args) {
        Solution s = new Solution();
        ListNode []test_cases = {
            new ListNode(4, new ListNode(2, new ListNode(1, new ListNode(3)))),
            new ListNode(-1, new ListNode(5, new ListNode(3, new ListNode(4, new ListNode(0)))))
        };
        for (ListNode test_case : test_cases) {
            util.pr(test_case);
            util.pr(s.insertionSortList(test_case));
        }
    }
}