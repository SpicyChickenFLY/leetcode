# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right

def printTree(root: TreeNode):
    print()

class Solution:

    def recursion(self, depth: int, root: TreeNode) -> int:
        if root == None:
            return depth
        return max(depth + 1, self.recursion(depth + 1, root.left), self.recursion(depth+1, root.right))
    
    # RunTime: 43.99%, MemoryUsage: 12.50%
    def maxDepth(self, root: TreeNode) -> int:
        return self.recursion(0, root)
        

if __name__ == "__main__":
    s = Solution()
    test_cases = [
        TreeNode(1, TreeNode(2, TreeNode(3), TreeNode(4)), TreeNode(2, TreeNode(4), TreeNode(3))),
        TreeNode(1, TreeNode(2, None, TreeNode(3)), TreeNode(2, None, TreeNode(3, TreeNode(1)))),
        TreeNode(1, TreeNode(2, TreeNode(3)), TreeNode(2, TreeNode(3))),
        TreeNode(3, TreeNode(9), TreeNode(20, TreeNode(15), TreeNode(7)))
    ]
    for test_case in test_cases:
        print(s.maxDepth(test_case))
