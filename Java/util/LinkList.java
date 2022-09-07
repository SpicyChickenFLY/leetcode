class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; next = null; }
    ListNode(int x, ListNode next) { val = x; this.next = next; }
}

class util {
    class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; next = null; }
        ListNode(int x, ListNode next) { val = x; this.next = next; }
    }

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
}

