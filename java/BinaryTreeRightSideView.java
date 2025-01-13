import java.util.*;

class TreeNode {
    int val;
    TreeNode left, right;

    TreeNode(int x) {
        val = x;
        left = null;
        right = null;
    }

    TreeNode(int x, TreeNode left, TreeNode right) {
        val = x;
        this.left = left;
        this.right = right;
    }
}

class Solution {
    private List<Integer> result;
    // Iterative solution - Run time: 100.00%, Memory Usage: 100.00%
    public void recursion(TreeNode root, int depth) {
        if (root == null)
            return;
        if (result.size() <= depth)
            result.add(root.val);
        else
            result.set(depth, root.val);
        recursion(root.left, depth + 1);
        recursion(root.right, depth + 1);
        return;
    }

    public List<Integer> rightSideView(TreeNode root) {
        result = new ArrayList<Integer> ();
        recursion(root, 0);
        return result;
    }
}

public class BinaryTreeRightSideView {
    public static void main(String[] args) {
        // null; // new TreeNode(1);
        Solution s = new Solution();
        TreeNode root = new TreeNode(1, new TreeNode(2, null, new TreeNode(5)), new TreeNode(3));
        List<Integer> result = s.rightSideView(root);
        for (int var : result) {
            System.out.print(var);
        }
    }
}