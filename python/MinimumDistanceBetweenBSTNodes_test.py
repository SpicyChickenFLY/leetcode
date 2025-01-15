from typing import List, Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def getMinimumDifference(self, root: Optional[TreeNode]) -> int:
        if root == None:
            return 0
        min_diff = float("inf")
        last_val = -1
        while root != None:
            if root.left == None:
                # 不存在左子树，向右遍历
                if last_val != -1:
                    min_diff = min(min_diff, root.val - last_val)
                last_val = root.val

                root = root.right
            else:
                # 找当前节点在左子树中的前驱节点
                p_s = root.left # proceedor searcher
                while p_s.right != None and p_s.right != root:
                    p_s = p_s.right
                # 此时p_s就是前驱节点

                if p_s.right != root:
                    # 未建立当前节点和前驱节点的连接
                    p_s.right = root # 创建连接
                    root = root.left  # 继续寻找左子树节点驱动
                else: # p_s.right == root
                    # 已经建立当前节点和前驱节点的连接，即，左子树已遍历完成
                    p_s.right = None
                    if last_val != -1:
                        min_diff = min(min_diff, root.val - last_val)
                    last_val = root.val
                    root = root.right # 向后驱动
        return int(min_diff)

def test_all():
    s = Solution()
    test_cases = [
            {"input": [TreeNode(1, TreeNode(0), TreeNode(48, TreeNode(12), TreeNode(49)))], "expect": 1},
    ]
    for case in test_cases:
        assert s.getMinimumDifference(*case["input"]) == case["expect"]

if __name__ == '__main__':
    test_all()

