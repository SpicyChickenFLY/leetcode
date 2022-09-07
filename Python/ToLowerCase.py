class Solution:
    # Runtime: 57.38%  Memory Usage: 5.52%
    def toLowerCase(self, s) -> str:
        return "".join([chr(ord(i) + 32) if i >= "A" and i <= "Z" else i for i in s])

if __name__ == "__main__":
    test_cases = [
        "aAaA",
        "AbcD",
        "al&phaBET"
    ]
    s = Solution()
    for test_case in test_cases:
        print(s.toLowerCase(test_case))