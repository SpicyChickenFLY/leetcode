class TreeNode:    
    def __init__(self, x, left = None, right = None):
        self.val = x
        self.left = left
        self.right = right

class Solution:
    def compare(self):
        pass

    def recursion(self, root):
        if root == None:
            return [True, None, None]
        if root.left == None and root.right == None:
            return [True, root.val, root.val]
        l_result = self.recursion(root.left)
        if not l_result[0]:
            return [False, None, None] 
        r_result = self.recursion(root.right)
        if not r_result[0]:
            return [False, None, None] 
        if (l_result[1] == None or l_result[1] < root.val) and \
            (r_result[2] == None or root.val < r_result[2]):
            max_v = r_result[1] if r_result[1] != None else root.val
            min_v = l_result[2] if l_result[2] != None else root.val
            return [True, max_v, min_v]
        else:
            return [False, None, None] 
        
    # runtime: 77.24% memory usage: 26.76%
    def isValidBST(self, root):
        return self.recursion(root)[0]

if __name__ == "__main__":
    inputs = [
        TreeNode(5, TreeNode(14, TreeNode(1), None), None),
        TreeNode(3, TreeNode(1, TreeNode(0), TreeNode(2, None, TreeNode(3))), TreeNode(5, TreeNode(4), TreeNode(6)))
    ]
    s = Solution()
    for i in inputs:
        print(s.isValidBST(i))