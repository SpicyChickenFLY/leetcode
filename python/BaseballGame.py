class Solution:
    # My solution - Runtime: 92.41% MemoryUsage: 37.01%%
    def calPoints(self, ops) -> int:
        stack = []
        for op in ops:
            if op == "C":
                stack.pop()
            elif op == "D":
                stack.append(stack[-1] * 2)
            elif op == "+":
                stack.append(stack[-1] + stack[-2])
            else:
                stack.append(int(op))
        return sum(stack)

if __name__ == "__main__":
    test_cases= [
        ["5","2","C","D","+"],
        ["5","-2","4","C","D","9","+","+"]
    ]
    s = Solution()
    for test_case in test_cases:
        print(s.calPoints(test_case))