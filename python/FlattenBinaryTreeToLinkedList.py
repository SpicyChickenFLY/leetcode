# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right

    def printTree(self):
        result = []
        self.traverseLevel(0, result, self)
        print(result)

    def traverseLevel(self, depth, result, root):
        while depth > len(result) - 1:
            result.append([])
        if root != None:
            result[depth].append(root.val)
            self.traverseLevel(depth+1, result, root.left)
            self.traverseLevel(depth+1, result, root.right)
        else:
            result[depth].append(None)

class Solution:
    # RunTime: 84.13%, MemoryUsage: 8.70%
    def flatten_sub(self, root: TreeNode):
        if root == None:
            return None

        left_end = self.flatten_sub(root.left)
        right_end = self.flatten_sub(root.right)

        if left_end == None and right_end == None:
            return root

        if left_end != None: 
            left_end.right = root.right
            root.right = root.left
            root.left = None

        return right_end if right_end != None else left_end
 

    def flatten(self, root: TreeNode) -> None:
        self.flatten_sub(root)

    
if __name__ == "__main__":
    test_cases = [
        TreeNode(1, TreeNode(2, TreeNode(3), TreeNode(4)), TreeNode(5, None, TreeNode(6)))
    ]
    s = Solution()
    for test_case in test_cases:
        test_case.printTree()
        s.flatten(test_case)
        test_case.printTree()