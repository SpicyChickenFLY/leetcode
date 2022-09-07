# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right

class Solution:
    def recursion(self, root: TreeNode, val):
        val = val * 10 + root.val
        if root.left == None and root.right == None:
            return val
        elif root.left == None:
            return self.recursion(root.right, val)
        elif root.right == None:
            return self.recursion(root.left, val)
        else:
            return self.recursion(root.left,val) + self.recursion(root.right, val)

    # RunTime: 16.48%%, MemoryUsage: 5.55%
    def sumNumbers(self, root: TreeNode) -> int:
        if root == None:
            return 0
        else:
            return self.recursion(root, 0)

if __name__ == "__main__":
    s = Solution()
    test_cases = [
        TreeNode(1, TreeNode(2), TreeNode(3)),
        TreeNode(4, TreeNode(9, TreeNode(5), TreeNode(1)), TreeNode(0))
    ]
    for test_case in test_cases:
        print(s.sumNumbers(test_case))