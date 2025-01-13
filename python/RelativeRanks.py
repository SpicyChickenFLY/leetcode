class Solution:
    # Runtime: 30.89%  Memory Usage: 5.69%
    def findRelativeRanks(self, nums: List[int]) -> List[str]:
        sorted_nums = sorted(nums, reverse=True)
        result = {}

        for i in range(0, len(nums)):
            if i == 0:
                result[sorted_nums[i]] = "Gold Medal"
            elif i == 1:
                result[sorted_nums[i]] = "Silver Medal"
            elif i == 2:
                result[sorted_nums[i]] = "Bronze Medal"
            else:
                result[sorted_nums[i]] = str(i+1)

        return list(map(lambda x: result[x], nums))


if __name__ == "__main__":
    s = Solution()
    test_cases = [
        [5, 4, 3, 2, 1]
    ]
    for test_case in test_cases:
        print(test_case, ": ", s.findRelativeRanks(test_case))
