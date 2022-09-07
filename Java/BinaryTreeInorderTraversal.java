import java.util.*;

class TreeNode {
    int val;
    TreeNode left, right;

    TreeNode(int x) {
        val = x;
        left = null;
        right = null;
    }
}

class Solution {
    // Iterative solution - Run time: 51.91%, Memory Usage: 99.99%
    public List<Integer> inorderTraversal1(TreeNode root) {
        List<Integer> result = new ArrayList<Integer>();
        Stack<TreeNode> stack = new Stack<TreeNode>();
        Stack<Integer> middle = new Stack<Integer>();
        if (root == null)
            return result;
        TreeNode p = root;
        while (true) {
            if (p == null) {
                if (!stack.isEmpty()) {
                    result.add(middle.pop());
                    p = stack.pop();
                } else {
                    break;
                }
            } else if (p.left != null) {
                middle.push(p.val);
                stack.push(p.right);
                p = p.left;
            } else if (p.right != null) {
                result.add(p.val);
                p = p.right;
            } else {
                result.add(p.val);
                p = null;
            }  
        }
        return result;
    }

    // Iterative solution - Run time: 51.91%, Memory Usage: 99.99%
    public List<Integer> inorderTraversal(TreeNode root) {
        if (root == null)
            return new ArrayList<Integer>();
        List<Integer> result = inorderTraversal(root.left),
        temp = inorderTraversal(root.right);
        result.add(root.val);
        for (int var : temp) {
            result.add(var);
        }
        return result;
    }
}

public class BinaryTreeInorderTraversal {
    public static void main(String[] args) {
        // null; // new TreeNode(1);
        TreeNode root       =   new TreeNode(1);
        root.left           =   new TreeNode(4);
        root.right          =   new TreeNode(3);
        root.left.left      =   new TreeNode(2);
        root.left.right     =   null;
        root.right.left     =   null;
        root.right.right    =   null;
        Solution s = new Solution();
        List<Integer> result = s.inorderTraversal(root);
        for (int var : result) {
            System.out.println(var);
        }
    }
}