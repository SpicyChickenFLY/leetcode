
// Definition for a binary tree node. 
class TreeNode { 
    int val; 
    TreeNode left; 
    TreeNode right; 
    TreeNode(int x) { val = x; left = null; right = null; } 
    TreeNode(int x, TreeNode left, TreeNode right) {val = x; this.left = left; this.right = right;}
}

class Solution {
    // Runtime: 12.72%, Memory Usage: 100.00%
    public int minDepth(TreeNode root) {
        if (root == null) return 0;
        if (root.left == null && root.right == null)  return 1;
        if (root.left == null) return minDepth(root.right) + 1;
        if (root.right == null) return minDepth(root.left) + 1;
        int left = minDepth(root.left);
        int right = minDepth(root.right);
        System.out.println("root:" + root.val + ", left:" + left + ", right:" + right);
        return ( left < right ? left : right ) + 1;
    }

    // Best: 100.00%
    public int minDepth1(TreeNode root) {
        if (root == null) {
            return 0;
        }
        if (root.left == null) {
            return 1 + minDepth(root.right);
        }
        if (root.right == null) {
            return 1 + minDepth(root.left);
        }
        return 1 + Integer.min(minDepth(root.left), minDepth(root.right));
    }
}

public class MinimumDepthOfBinaryTree {
    public static void main(String[] args) {
        Solution s = new Solution();
        TreeNode []testCases = {
            new TreeNode(1, new TreeNode(2), null),
            new TreeNode(3, new TreeNode(9), new TreeNode(20, new TreeNode(15), new TreeNode(7)))
        };
        for (TreeNode testCase : testCases) {
            System.out.println(s.minDepth(testCase));
        }
    }
}