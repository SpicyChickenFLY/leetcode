class Solution:
    
    def longestValidParentheses(self, s):
        s = list(s)
        length = len(s)
        while True:
            for index in range(len(s)):
                if index  + 1 < len(s) and s[index] == '(' and s[index + 1] == ')':
                    del s[index]
                    del s[index]
                    s.insert(index, 2)
                elif index + 2 < len(s) and s[index] == '(' and type(s[index + 1]) == int and s[index + 2] == ')':
                    s[index + 1] += 2
                    del s[index]
                    del s[index + 1]
                elif index  + 1 < len(s) and type(s[index]) == int and type(s[index + 1]) == int:
                    s[index] += s[index + 1]
                    del s[index + 1]
            if len(s) == length:
                break
            else:
                length = len(s)
        max_value = 0
        for c in s:
            if type(c) == int and c > max_value:
                max_value = c
        return max_value

if __name__ == "__main__":
    s = Solution()
    inputs = [
        ')()(())',
        ')()())',
        '(()'
    ]
    for i in inputs: 
        print(s.longestValidParentheses(i))