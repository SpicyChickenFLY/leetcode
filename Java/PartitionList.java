import com.sun.scenario.effect.impl.state.LinearConvolveRenderState.PassType;

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
        next = null;
    }

    ListNode(int x, ListNode n) {
        val = x;
        next = n;
    }

    void pr() {
        System.out.print(val + "->");
        ListNode temp = next;
        while (temp != null) {
            System.out.print(temp.val + "->");
            temp = temp.next;
        }
        System.out.println("NULL");
    }
}



class Solution {
    // Run time: 100.00.%, Memory Usage: 100.00%
    public ListNode partition(ListNode head, int x) {
        ListNode middle = new ListNode(x), p_1 = head, p_2 = middle;
        if (head == null || head.next == null)
            return head;
        while (p_1.next != null) {
            if (p_1.next.val < x) {
                p_1 = p_1.next;
            } else {
                p_2.next = p_1.next;
                p_1.next = p_1.next.next;
                p_2.next.next = null;
                p_2 = p_2.next;
            } 
        }
        if (head.val  >= x && head != p_1) {
            middle.val = head.val;
            head = head.next;                
        } else
            middle = middle.next;
        p_1.next = middle;
        return head;
    }
}

public class PartitionList {
    public static void main(String[] args) {
        ListNode[] lists = {
            new ListNode(1, new ListNode(4, new ListNode(3, new ListNode(2, new ListNode(5, new ListNode(2)))))),
            new ListNode(1),
            new ListNode(1, new ListNode(1))
        };
        int[] x = {
            3,
            0,
            0
        };
        Solution s = new Solution();
        for (int i = 0; i < x.length; i++) {
            lists[i].pr();
            ListNode result = s.partition(lists[i], x[i]);
            result.pr();
        }
    }
}