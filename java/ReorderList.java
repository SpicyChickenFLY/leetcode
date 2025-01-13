// Definition for singly-linked list.
class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; next = null; }
    ListNode(int x, ListNode next) { val = x; this.next = next; }
}

class Solution {
    public void printList(ListNode head) {
        while (head != null) {
            System.out.print(head.val + "->");
            head = head.next;
        }
        System.out.println("null");
    }

    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null)
            return head;
        ListNode newHead = reverseList(head.next);
        head.next.next = head;
        head.next = null;
        return newHead;
    }

    // Run time: 99.95%, Memory Usage: 97.73%
    public void reorderList(ListNode head) {
        if (head == null)
            return;
        if (head.next == null)
            return;   
        ListNode sp = head, fp = head;
        ListNode temp1, temp2;
        while (true) {
            if (fp.next == null)
                break;
            else if (fp.next.next == null)
                break;
            fp = fp.next.next;
            sp = sp.next;
        }
        fp = sp.next;
        sp.next = null;
        fp = reverseList(fp);
        sp = head;
        
        while (fp != null) {
            temp1 = fp;
            fp = fp.next;

            temp2 = sp.next;
            sp.next = temp1;
            temp1.next = temp2;
            sp = sp.next.next;
        }
    }
}

public class ReorderList {
    public static void main(String[] args) {
        Solution s = new Solution();
        ListNode []testCases = {
            null,
            new ListNode(1, new ListNode(2, new ListNode(3))),
            new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4)))),

        };
        for (ListNode testCase : testCases) {
            s.printList(testCase);
            s.reorderList(testCase);
            s.printList(testCase);
        }
    }
}
