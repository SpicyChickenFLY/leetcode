class Solution:
    # Runtime: 12.25% MemoryUsage: 5.13%
    def fib(self, N: int) -> int:
        if N == 0:
            return 0
        if N == 1:
            return 1
        return self.fib(N - 1) + self.fib(N - 2)

if __name__ == "__main__":
    s = Solution()
    for test_case in range(20, 0, -1):
        print(str(test_case) + ":" + str(s.fib(test_case)))