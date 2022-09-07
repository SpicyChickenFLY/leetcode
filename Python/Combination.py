import copy


class Solution(object):
    # Runtime: 79.56%, MemoryUsage: 8.94%
    def combine(self, n, k):
        if k > n or k == 0:
            return []
        if k == 1:
            return [[i] for i in range(1, n+1)]
        if k == n:
            return [[i for i in range(1, n+1)]]

        former_result = self.combine(n-1, k)
        for item in self.combine(n-1, k-1):
            item.append(n)
            former_result.append(item)

        return former_result

if __name__ == "__main__":
    s = Solution()
    for k in range(1, 5):
        for n in range(k, 5):
            print("{0}, {1}, {2}" .format(n, k, s.combine(n, k)))
