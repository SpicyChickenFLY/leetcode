class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; next = null; }
    ListNode(int x, ListNode next) { val = x; this.next = next; }
}

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x, TreeNode left, TreeNode right) {
        val = x;
        this.left = left;
        this.right = right;
    }

    TreeNode(int x) {
        val = x;
        left = null;
        right = null;
    }
}

class Solution {
    // Run time: 100.00%, Memory Usage: 100.00%
    public TreeNode sortedListToBST(ListNode head) {
        if (head == null) return null;
        if (head.next == null) return new TreeNode(head.val);
        ListNode fp = head, sp = head, l = head;
        while (true) {
            if (fp.next == null)
                break;
            else if (fp.next.next == null)
                break;
            fp = fp.next.next;
            l = sp;
            sp = sp.next;
        }

        TreeNode root = new TreeNode(sp.val);
        fp = sp.next;
        l.next = null;
        if (l == sp) head = null;
        root.left = sortedListToBST(head);
        root.right = sortedListToBST(fp);

        return root;
    }
}

public class ConvertSortedListToBinarySearchTree {
    public static void main(String[] args) {
        Solution s = new Solution();
        ListNode []testCases = {
            null,
            new ListNode(0),
            new ListNode(-10, new ListNode(-3, new ListNode(0, new ListNode(5, new ListNode(9)))))
        };
        for (ListNode testCase : testCases) {
            s.sortedListToBST(testCase);
        }
    }
}