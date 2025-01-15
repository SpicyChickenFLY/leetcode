from typing import List, Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
        if root == None:
            return None

        # 二叉搜索节点
        current = root
        prev = None
        while current != None:
            if key == current.val:
                break
            prev = current
            current = current.right if key > current.val else current.left
        if current == None:
            # 搜索失败
            return root

        sub = None
        if current.left != None:
            # 存在左子树
            temp = current.left
            while temp.right != None:
                temp = temp.right
            temp.right = current.right
            sub = current.left
        elif current.right != None:
            # 存在右子树
            temp = current.right
            while temp.left != None:
                temp = temp.left
            temp.left = current.left
            sub = current.right

        if prev == None:
            root = sub
        elif key > prev.val:
            prev.right = sub
        else:
            prev.left = sub

        return root

def test_all():
    s = Solution()
    test_cases = [
            {
                "input": [ TreeNode(5, TreeNode(2, None, TreeNode(4)), TreeNode(6, None, TreeNode(7))), 0 ],
                 "expect": TreeNode(5, TreeNode(2, None, TreeNode(4)), TreeNode(6, None, TreeNode(7)))
            },
    ]
    for case in test_cases:
        # assert s.deleteNode(*case["input"]) == case["expect"]
        print("\n")


if __name__ == '__main__':
    test_all()

