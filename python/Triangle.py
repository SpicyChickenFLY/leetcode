class Solution:
    # Runtime: 78.53% MemoryUsage: 6.67%
    def minimumTotal(self, triangle) -> int:
        if len(triangle) == 0:
            return 0
        for layer in range(1, len(triangle)):
            triangle[layer][0] += triangle[layer-1][0]
            triangle[layer][-1] += triangle[layer-1][-1]
            for index in range(1, len(triangle[layer]) - 1):
                triangle[layer][index] += min(triangle[layer-1][index-1], triangle[layer-1][index])
        print(triangle)
        return min(triangle[-1])

if __name__ == "__main__":
    test_cases = [
        [
            [2],
            [3,4],
            [6,5,7],
            [4,1,8,3]
        ]
    ]
    s = Solution()
    for test_case in test_cases:
        print(s.minimumTotal(test_case))