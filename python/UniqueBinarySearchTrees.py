class Solution:
    def __init__(self):
        self.table = [1, 1, 2, 5]

    # Runtime: 68.86%, Memory Usage: 5.78%
    def recursion(self, n):
        result = 0
        for i in range(n - 1, -1, -1):
            result += self.table[i] * self.table[n - 1 - i]
        self.table.append(result)

    def numTrees(self, n):
        for i in range(len(self.table), n + 1):
            self.recursion(i)
        return self.table[n]

if __name__ == "__main__":
    s = Solution()
    for n in range(7, 2, -1):
        print(s.numTrees(n))
    print(s.table)
