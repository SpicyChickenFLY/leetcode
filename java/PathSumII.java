import java.util.*;

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
    List<List<Integer>> result;

    public void recursion(TreeNode root, int sum, List<Integer> temp) {
        if (root == null)
            return;
        //System.out.printf("%d:%d", root.val, sum);
        temp.add(root.val);
        if (root.left == null && root.right == null && sum == root.val) {
            List<Integer> newList = new ArrayList<Integer>();
            for (int i = 0; i < temp.size(); i++)
                newList.add(temp.get(i));
            result.add(newList);
        }
        recursion(root.left, sum - root.val, temp);
        recursion(root.right, sum - root.val, temp);
        temp.remove(temp.size() - 1);
    }

    // Runtime: 100.00%, Memory Usage: 100.00%
    public List<List<Integer>> pathSum(TreeNode root, int sum) {
        result = new ArrayList<List<Integer>>();
        List<Integer> temp = new ArrayList<Integer>();
        recursion(root, sum, temp);
        return result;
    }

};

public class PathSumII {
    public static void main(String[] args) {
        Solution s = new Solution();
        TreeNode test_case = null;
        List<List<Integer>> result = s.pathSum(test_case, 10);
        System.out.println(result.size());
        test_case = new TreeNode(1);
        result = s.pathSum(test_case, 10);
        System.out.println(result.size());
        test_case = new TreeNode(5, 
                new TreeNode(4, new TreeNode(11, new TreeNode(7), new TreeNode(2)), null), 
                new TreeNode(8, new TreeNode(13), new TreeNode(4, null, new TreeNode(1)))
                );
        result = s.pathSum(test_case, 22);
        System.out.println(result.size());
        test_case = new TreeNode(5, new TreeNode(4, new TreeNode(11, new TreeNode(7), new TreeNode(2)), null),
                new TreeNode(8, new TreeNode(13), new TreeNode(4, new TreeNode(5), new TreeNode(1))));
        result = s.pathSum(test_case, 22);
        System.out.println(result.size());
    }
};