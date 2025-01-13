# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right

class Solution:

    def recursion(self, depth: int, result, root: TreeNode | None) -> None:
        '''BFS,广度遍历,一层一层来'''
        if root is not None:
            while depth > len(result) - 1:
                result.append([])
            # 每层结果都记录在一个列表中
            result[depth].append(root.val)
            # 先左子树，再右子树，保证顺序
            self.recursion(depth+1, result, root.left)
            self.recursion(depth+1, result, root.right)

    # RunTime: 64.84%, MemoryUsage: 5.41%
    def zigzagLevelOrder(self, root: TreeNode):
        '''先层级遍历所有节点，随后根据层级反序'''
        result = []
        if root is not None:
            self.recursion(0, result, root)
        for level, level_result in enumerate(result):
            if level % 2 == 1:
                result[level] = list(reversed(level_result))
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
        print(s.zigzagLevelOrder(test_case))
