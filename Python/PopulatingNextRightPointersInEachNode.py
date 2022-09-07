
# Definition for a Node.
class Node:
    def __init__(self, val, left=None, right=None, next=None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next

class Solution:
    # RunTime: 5.69%, MemoryUsage: 7.14%
    def connect(self, root: 'Node') -> 'Node':
        if root == None or root.left == None:
            return root
        else:
            vertical = root
            while vertical.left != None:
                horizon = vertical
                while horizon != None:
                    horizon.left.next = horizon.right
                    if horizon.next != None:
                        horizon.right.next = horizon.next.left
                    horizon = horizon.next
                vertical = vertical.left
            return root

if __name__ == "__main__":
    test_case = Node(1, Node(2, Node(4), Node(5)), Node(3, Node(6), Node(7)))
    s = Solution()
    s.connect(test_case)
