# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right

def printTree(root: TreeNode):
    print()

class Solution:
    # RunTime: 40.94%, MemoryUsage: 5.17%
    def compare(self, left: TreeNode, right: TreeNode) -> bool:
        if left == None and right == None:
            return True
        return False if left == None or right == None or left.val != right.val or not self.compare(left.left, right.right) or not self.compare(left.right, right.left) else True

    def isSymmetric(self, root: TreeNode) -> bool:
        if root == None:
            return True
        return self.compare(root.left, root.right)

if __name__ == "__main__":
    s = Solution()
    test_cases = [
        TreeNode(1, TreeNode(2, TreeNode(3), TreeNode(4)), TreeNode(2, TreeNode(4), TreeNode(3))),
        TreeNode(1, TreeNode(2, None, TreeNode(3)), TreeNode(2, None, TreeNode(3))),
        TreeNode(1, TreeNode(2, TreeNode(3)), TreeNode(2, TreeNode(3)))
    ]
    for test_case in test_cases:
        print(s.isSymmetric(test_case), "\n")
