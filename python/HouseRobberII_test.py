from typing import List


class Solution:
    def rob(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        if len(nums) < 4:
            return max(nums)
        return max(self._rob(nums[1:]), self._rob(nums[:-1]))

    def _rob(self, nums: List[int]) -> int:
        dp = [0] * len(nums)
        dp[0] = nums[0]
        dp[1] = max(dp[0], nums[1])
        for i in range(2, len(nums)):
            dp[i] = max(dp[i-1], dp[i-2] + nums[i])
        return dp[-1]

    def __init__(self) -> None:
        pass

def test_all():
    s = Solution()
    test_cases = [
            {"input": [[6,6,4,8,4,3,3,10]], "expect": 27},
            {"input": [[2, 1, 1, 2]], "expect": 3},
            {"input": [[2, 3, 2]], "expect": 3},
            {"input": [[1, 2, 3, 1]], "expect": 4},
            {"input": [[1, 2, 3]], "expect": 3},
            {"input": [[3, 1, 4, 1]], "expect": 7},
            {"input": [[4, 1, 2, 7, 5, 3, 1]], "expect": 14},
    ]
    for case in test_cases:
        print(case["input"])
        assert s.rob(*case["input"]) == case["expect"]
        print("\n")


if __name__ == '__main__':
    test_all()

