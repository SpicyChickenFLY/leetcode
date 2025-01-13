class Solution:
    def minimumLength(self, s: str) -> int:
        count_map = {}
        for c in s:
            if c not in count_map:
                count_map[c] = 1
            else:
                count_map[c] += 1
        length = 0
        for _, value in count_map.items():
            if value != 0:
                length += 2 - (value & 1)
        return length

if __name__ == "__main__":
    inputs = [
            "abaacbcbb",
            "aa"
    ]
    solution = Solution()
    for i in inputs:
        print(solution.minimumLength(i))
