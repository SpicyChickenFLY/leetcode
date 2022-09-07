class Solution:
    # Runtime: 95.59%  Memory Usage: 100.00%
    def defangIPaddr(self, address) -> str:
        return "[.]".join(address.split("."))

if __name__ == "__main__":
    test_cases = [
        "1.1.1.1",
        "255.100.50.0"
    ]
    s = Solution()
    for test_case in test_cases:
        print(s.defangIPaddr(test_case))