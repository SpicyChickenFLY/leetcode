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
    // RunTime: 100.00%, MemoryUsage: 31.58%
    public ListNode sortList(ListNode head) {
        ListNode result = null, left = null, right = null, iter = null, temp = null;
        if (head == null)
            return null;
        iter = head.next;
        head.next = null;
        while (iter != null) {
            temp = iter;
            iter = iter.next;
            if (temp.val > head.val) {
                temp.next = right;
                right = temp;
            } else if (temp.val < head.val) {
                temp.next = left;
                left = temp;
            } else {
                temp.next = head;
                head = temp;
            }
        }
        iter = result;
        left = sortList(left); 
        right = sortList(right);

        if (left != null) {
            result = left;
            iter = result;
            while (iter.next != null)
                iter = iter.next;
            iter.next = head;
        } else {
            result = head;
            iter = result;
        }
        while (iter.next != null)
            iter = iter.next;
        iter.next = right;
        return result;
    }
}

public class SortList {
    public static void main(String[] args) {
        Solution s = new Solution();
        ListNode []test_cases = {
            new ListNode(4, new ListNode(2, new ListNode(1, new ListNode(3)))),
            new ListNode(-1, new ListNode(5, new ListNode(3, new ListNode(4, new ListNode(0)))))
        };
        for (ListNode test_case : test_cases) {
            util.pr(test_case);
            util.pr(s.sortList(test_case));
        }
    }
}