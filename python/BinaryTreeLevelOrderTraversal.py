# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right

def printTree(root: TreeNode):
    print()

class Solution:

    def recursion(self, depth: int, result, root: TreeNode) -> None:
        if root != None:
            while depth > len(result) - 1:
                result.append([])
            result[depth].append(root.val)
            self.recursion(depth+1, result, root.left)
            self.recursion(depth+1, result, root.right)
    
    # RunTime: 48.91%, MemoryUsage: 6.45%
    def levelOrder(self, root: TreeNode):
        result = []
        if root != None:
            self.recursion(0, result, root)
        return result
        

if __name__ == "__main__":
    s = Solution()
    test_cases = [
        TreeNode(1, TreeNode(2, TreeNode(3), TreeNode(4)), TreeNode(2, TreeNode(4), TreeNode(3))),
        TreeNode(1, TreeNode(2, None, TreeNode(3)), TreeNode(2, None, TreeNode(3))),
        TreeNode(1, TreeNode(2, TreeNode(3)), TreeNode(2, TreeNode(3))),
        TreeNode(3, TreeNode(9), TreeNode(20, TreeNode(15), TreeNode(7)))
    ]
    for test_case in test_cases:
        print(s.levelOrder(test_case))
