class Solution:
    # Runtime: 38.50%  Memory Usage: 5.05%
    def removeOuterParentheses(self, S) -> str:
        result = ""
        stack = []
        for char in S:
            if char == "(":
                stack.append(char)
                if len(stack) != 1:
                    result += "("
            else:
                if len(stack) != 1:
                    result += ")"
                stack.pop()
        return result

if __name__ == "__main__":
    test_cases = [
        "(()())(())",
        "(()())(())(()(()))",
        "()()"
    ]
    s = Solution()
    for test_case in test_cases:
        print(s.removeOuterParentheses(test_case))