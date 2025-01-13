class Solution:
    # Runtime: 74.01%  Memory Usage: 5.27%
    def numJewelsInStones(self, J: str, S: str) -> int:
        return sum([S.count(jewel) for jewel in J]) 


if __name__ == "__main__":
    test_cases = [
        ["aA", "aAAbbba"],
        ["z", "ZZ"]
    ]
    s = Solution()
    for test_case in test_cases:
        print(s.numJewelsInStones(test_case[0], test_case[1]))