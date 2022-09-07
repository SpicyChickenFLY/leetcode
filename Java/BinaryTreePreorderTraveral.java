import java.util.*;

public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x, TreeNode left, TreeNode right) {
        val = x;
        this.left = left;
        this.right = right;
    }
}

class Solution {
    // RunTime: 60.53%, MemoryUsage: 100.00%
    public List<Integer> preorderTraversal(TreeNode root) {
        List<Integer> result = new ArrayList<Integer>();
        if (root == null)
            return result;
        result.add(root.val);
        result.addAll(preorderTraversal(root.left));
        result.addAll(preorderTraversal(root.right));
        return result;
    }
}

public class BinaryTreePreorderTraveral {
    public static void main(String[] args) {
        Solution s = new Solution();
        TreeNode input = new TreeNode(1, null, new TreeNode(2, new TreeNode(3, null, null), null));
        s.preorderTraversal(input);
    }
}