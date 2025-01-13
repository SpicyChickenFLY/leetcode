class Solution:
    def __init__(self):
        self.result = [
            [1],
            [1, 1],
            [1, 2, 1]
        ]

    # Runtime: 41.51% MemoryUsage: 7.14%
    def generate(self, numRows):
        if len(self.result) >= numRows:
            return self.result[:numRows]
        for layer in range(len(self.result), numRows):
            self.result.append([1] * (len(self.result) + 1))
            for index in range(1, len(self.result[layer]) - 1):
                self.result[layer][index] = self.result[layer-1][index-1] + self.result[layer-1][index]
        return self.result

if __name__ == "__main__":
    s = Solution()
    print(s.generate(5))
