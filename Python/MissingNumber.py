class Solution:
    # Runtime - 99.45%, Memory Usage - 27.06%
    def missingNumber(self, nums):
        n = len(nums)
        return (n + 1) * n // 2 - sum(nums)

if __name__ == "__main__":
    inputs = [
        [3, 0, 1],
        [0, 1, 2, 3],
        [9, 6, 4, 2, 3, 5, 7, 0, 1]
    ]
    s = Solution()
    for i in inputs:
        print(s.missingNumber(i))