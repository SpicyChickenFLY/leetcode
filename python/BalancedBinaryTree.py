# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right

class Solution:
    def distance(self, root):
        if root == None:
            return 0
        left = self.distance(root.left)
        right = self.distance(root.right)
        if left == -1 or right == -1:
            return -1
        d = abs(left - right)
        # print(root.val, left, right, d)
        return -1 if d > 1 else max(left, right) + 1

    # RunTime: 51.66%, MemoryUsage: 54.29%
    def isBalanced(self, root: TreeNode) -> bool:
        return False if self.distance(root) == -1 else True


if __name__ == "__main__":
    test_cases = [
        TreeNode(3, TreeNode(9), TreeNode(20, TreeNode(15), TreeNode(7))),
        TreeNode(1, TreeNode(2, TreeNode(3, TreeNode(4), TreeNode(4)), TreeNode(3)), TreeNode(2)),
    ]

    s = Solution()
    for test_case in test_cases:
        print(s.isBalanced(test_case))
