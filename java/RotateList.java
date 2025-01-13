class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
        next = null;
    }
}

class Solution {    

    public void printList(ListNode head) {
        ListNode temp = head;
        while (temp != null) {
            System.out.print(temp.val + " ");
            temp = temp.next;
        }
        System.out.println("NULL");
    }

    public ListNode createList(int arr[]) {
        ListNode head = new ListNode(arr[0]), temp = head;
        for (int i = 1; i < arr.length; i++) {
            temp.next = new ListNode(arr[i]);
            temp = temp.next;
        }
        return head;
    }

    public ListNode rotateRight(ListNode head, int k) {
        if (k == 0 || head == null) {
            return head;
        }else {
            ListNode temp = head;
            int length = 1;
            while (temp.next != null) {
                length++;
                temp = temp.next;
            }
            temp.next = head;
            k = length - (k % length);
            for (int i = 1; i < k; i++) {
                head = head.next;
            }
            ListNode newHead = head.next;
            head.next = null;
            return newHead;
        }
    }
}

public class RotateList { 
    public static void main(String[] args) {
        Solution s = new Solution();
        int arr[] = { 0, 1, 2 };
        ListNode head = s.createList(arr);
        for (int i = 0; i < 10; i++) {
            head = s.rotateRight(head, i);
            s.printList(head);
        }  
    }
}