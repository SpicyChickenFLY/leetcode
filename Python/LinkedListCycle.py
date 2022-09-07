# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x, next=None):
        self.val = x
        self.next = next


class Solution(object):
    # RunTime: 5.05%, MemoryUsage: 58.51%
    def hasCycle1(self, head):
        result = []
        while head != None:
            if head.next in result:
                return True
            else:
                result.append(head.next)
                head = head.next
        return False

    # RunTime: 63.28%, MemoryUsage: 52.17%
    def hasCycle(self, head):
        while head != None:
            if head.val == None:
                return True
            else:
                head.val = None
                head = head.next
        return False


if __name__ == "__main__":
    s = Solution()
    test_cases = [
        [ListNode(3, ListNode(2, ListNode(0, ListNode(2)))), 1],
        [ListNode(1, ListNode(2)), 0],
        [ListNode(1, ListNode(2)), -1]
    ]
    for test_case in test_cases:
        head = test_case[0]
        tail = head
        while tail.next != None:
            tail = tail.next
        if test_case[1] >= 0:
            temp = head
            for i in range(test_case[1]):
                temp = temp.next
            tail.next = temp
        print(s.hasCycle(head))
