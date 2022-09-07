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
    // Runtime: 100.00%, Memory Usage: 100.00%
    class Record {
        public int index;
        public int data;
    }
    public int kthSmallest(TreeNode root, int k) {
        Record r = new Record();
        kthSmallest(root, k, r);
        return r.data;
    }
    
    private void kthSmallest(TreeNode root, int k, Record r) {
        if (root == null)
            return;
        kthSmallest(root.left, k, r);
        r.index ++;
        if (r.index == k) {
            r.data = root.val;
            return;
        }
        kthSmallest(root.right, k, r);
    }
}