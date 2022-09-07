class Solution:
    def __init__(self):
        self.result = [
            [1],
            [1, 1],
            [1, 2, 1]
        ]

    # Runtime: 94.46% MemoryUsage: 7.69%
    def getRow(self, rowIndex):
        if len(self.result) > rowIndex:
            return self.result[rowIndex]
        for layer in range(len(self.result), rowIndex + 1):
            self.result.append([1] * (len(self.result) + 1))
            for index in range(1, len(self.result[layer]) - 1):
                self.result[layer][index] = self.result[layer-1][index-1] + self.result[layer-1][index]
        return self.result[-1]

if __name__ == "__main__":
    s = Solution()
    print(s.getRow(5))
    print(s.getRow(3))