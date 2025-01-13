import java.util.*;

class TreeNode { 
    int val; 
    TreeNode left; 
    TreeNode right; 
    TreeNode(int x) { 
        val = x; 
    }
    TreeNode(int x, TreeNode left, TreeNode right) {
        val = x;
        this.left = left;
        this.right = right;
    }
}

//sort solution - Run time: 5.94%, Memory Usage: 85.18%
class BSTIterator {
    private List<Integer> iterator;

    public BSTIterator(TreeNode root) {
        iterator = inorderTraversal(root);
    }

    public List<Integer> inorderTraversal(TreeNode root) {
        if (root == null)
            return new ArrayList<Integer>();
        List<Integer> result = inorderTraversal(root.left);
        result.add(root.val);
        for (int var : inorderTraversal(root.right))
            result.add(var);
        return result;
    }

    public int next() {
        int temp = iterator.get(0);
        iterator.remove(0);
        return temp;
    }

    public boolean hasNext() {
        return iterator.size() > 0;
    }
}

public class BinarySearchTreeIterator {
    public static void main(String[] args) {
        TreeNode bst = new TreeNode(7, new TreeNode(3), new TreeNode(15, new TreeNode(9), new TreeNode(20)));
        BSTIterator obj = new BSTIterator(bst);
        while(obj.hasNext()) 
            System.out.println(obj.next());
    }
}


