class TreeNode {
    int val;
    TreeNode left, right;

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
};

class Solution {
    public boolean recursion(TreeNode root, int sum) {
        if (root == null)
            return false;
        //System.out.printf("%d:%d", root.val, sum);
        if (root.left == null && root.right == null) {
            return sum == root.val;
        }
        if (recursion(root.left, sum - root.val) || recursion(root.right, sum - root.val))
            return true;
        return false;
    }

    // Runtime: 100.00%, Memory Usage: 100.00%
    public boolean hasPathSum(TreeNode root, int sum) {
        return recursion(root, sum);
    }
};

public class PathSum {
    public static void main(String[] args) {
        Solution s = new Solution();
        TreeNode test_case = null;
        System.out.println(s.hasPathSum(test_case, 10));
        test_case = new TreeNode(1);
        System.out.println(s.hasPathSum(test_case, 10));
        test_case = new TreeNode(5, 
                new TreeNode(4, new TreeNode(11, new TreeNode(7), new TreeNode(2)), null), 
                new TreeNode(8, new TreeNode(13), new TreeNode(4, null, new TreeNode(1)))
                );
        System.out.println(s.hasPathSum(test_case, 22));
    }
}