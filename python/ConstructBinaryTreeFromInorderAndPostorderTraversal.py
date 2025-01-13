# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right

    def printTree(self):
        print(self.val)
        print("left: ", end='')
        if self.left != None:
            print(self.left.val)
            self.left.printTree()
        else:
            print(None)
        print("right: ", end='')
        if self.right != None:
            print(self.right.val)
            self.right.printTree()
        else:
            print(None)

class Solution:
    # RunTime: 34.61%, MemoryUsage: 11.11%
    def buildTree(self, inorder, postorder) -> TreeNode:
        print(inorder, postorder)
        root = None
        if len(postorder) > 0:
            peak = postorder.pop()
            root = TreeNode(peak)
            partition = inorder.index(peak)
            if partition > 0:
                left_inorder = inorder[:partition]
                root.left = self.buildTree(
                    left_inorder,
                    postorder[:partition]
                    )
            if partition < len(inorder) - 1:
                right_inorder = inorder[partition + 1:]
                root.right = self.buildTree(
                    right_inorder,
                    postorder[partition:]
                    )
        root.printTree()
        return root
        

if __name__ == "__main__":
    s = Solution()
    test_cases = [
        (
            [9, 3, 15, 20, 7],
            [9, 15, 7, 20, 3]
        ),
    ]
    for test_case in test_cases:
        result = s.buildTree(test_case[0], test_case[1])
        result.printTree()
