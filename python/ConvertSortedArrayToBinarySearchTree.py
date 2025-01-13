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
    # RunTime: 80.43%, MemoryUsage: 6.45%
    def sortedArrayToBST(self, nums):
        if len(nums) == 0:
            return None
        if len(nums) == 1:
            return TreeNode(nums[0])
        else:
            mid = len(nums) // 2
            return TreeNode(
                nums[mid], 
                self.sortedArrayToBST(nums[:mid]),
                self.sortedArrayToBST(nums[mid+1:])
                )

if __name__ == "__main__":
    test_cases = [
        [-10, -3, 0, 5, 9],
    ]
    s = Solution()
    for test_case in test_cases:
        result = s.sortedArrayToBST(test_case)
        result.printTree()