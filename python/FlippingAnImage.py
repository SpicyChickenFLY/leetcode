class Solution:
    # Runtime: 10.41%  Memory Usage: 5.27%
    def flipAndInvertImage(self, A):
        return [[1 if s==0 else 0 for s in l[::-1]] for l in A]

if __name__ == "__main__":
    test_cases = [
        [[1,1,0],[1,0,1],[0,0,0]],
        [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
    ]
    s = Solution()
    for test_case in test_cases:
        print(s.flipAndInvertImage(test_case))